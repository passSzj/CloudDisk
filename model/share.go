package model

import (
	"context"
	"fmt"
	"strconv"

	"github.com/google/uuid"
	"go-cloud-disk/cache"
	"go-cloud-disk/disk"
	"gorm.io/gorm"
)

type Share struct {
	Uuid        string `gorm:"primarykey"`
	Owner       string
	FileId      string // file uuid of share file
	FileName    string
	Title       string
	Size        int64
	SharingTime string
}

// SetEmptyShare set a empty share
func (share *Share) SetEmptyShare() {
	// remove share from dailyrank and add share into emptyshare set
	if share.DailyViewCount() > 10 {
		cache.RedisClient.ZRem(context.Background(), cache.DailyRankKey, share.Uuid)
		cache.RedisClient.SAdd(context.Background(), cache.EmptyShare, share.Uuid)
	}

	share.Owner = ""
	share.FileId = ""
	share.FileName = ""
	share.Title = "来晚了,分享的文件已被删除"
	share.Size = 0
	share.SharingTime = ""
}

// BeforeCreate create uuid before insert database
func (file *Share) BeforeCreate(tx *gorm.DB) (err error) {
	if file.Uuid == "" {
		file.Uuid = uuid.New().String()
	}
	return
}

// DownloadURL get share download url
func (share *Share) DownloadURL() (string, error) {
	var file File
	if err := DB.Where("uuid = ?", share.FileId).Find(&file).Error; err != nil {
		return "", fmt.Errorf("find user file err when build download url %v", err)
	}

	url, err := disk.BaseCloudDisk.GetObjectURL(file.FilePath, "", file.FileUuid+"."+file.FilePostfix)
	if err != nil {
		return "", fmt.Errorf("get object url err when get share download url, %v", err)
	}
	return url, nil
}

// ViewCount get share view from redis
func (share *Share) ViewCount() (num int64) {
	countStr, _ := cache.RedisClient.Get(context.Background(), cache.ShareKey(share.Uuid)).Result()
	if countStr == "" {
		return 0
	}
	num, _ = strconv.ParseInt(countStr, 10, 64)
	return
}

// DailyViewCount get daily view count by share uuid
func (share *Share) DailyViewCount() float64 {
	countStr := cache.RedisClient.ZScore(context.Background(), cache.DailyRankKey, share.Uuid).Val()
	return countStr
}

// AddViewCount add share view in redis
func (share *Share) AddViewCount() {
	cache.RedisClient.Incr(context.Background(), cache.ShareKey(share.Uuid))
	cache.RedisClient.ZIncrBy(context.Background(), cache.DailyRankKey, 1, share.Uuid)
}

// SaveShareInRedis save share info to redis
func (share *Share) SaveShareInfoToRedis(downloadUrl string) error {
	ctx := context.Background()
	// if Owner is not null, it means that a func has already
	// been written to redis
	if s := cache.RedisClient.HGet(ctx, cache.ShareInfoKey(share.Uuid), "Owner").Val(); s != "" {
		return nil
	}

	// use ptpeline to save share info to redis for ensure
	// share info all write to redis
	saveShare := cache.RedisClient.Pipeline()
	saveShare.HSet(ctx, cache.ShareInfoKey(share.Uuid), "Owner", share.Owner)
	saveShare.HSet(ctx, cache.ShareInfoKey(share.Uuid), "FileId", share.FileId)
	saveShare.HSet(ctx, cache.ShareInfoKey(share.Uuid), "FileName", share.FileName)
	saveShare.HSet(ctx, cache.ShareInfoKey(share.Uuid), "Title", share.Title)
	saveShare.HSet(ctx, cache.ShareInfoKey(share.Uuid), "Size", share.Size)
	saveShare.HSet(ctx, cache.ShareInfoKey(share.Uuid), "SharingTime", share.SharingTime)
	saveShare.HSet(ctx, cache.ShareInfoKey(share.Uuid), "downloadUrl", downloadUrl)
	_, err := saveShare.Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

// GetShareInfoFromRedis get share info from redis and return downloadurl
func (share *Share) GetShareInfoFromRedis() string {
	// if is a empty share fill empty message
	if cache.RedisClient.SIsMember(context.Background(), cache.EmptyShare, share.Uuid).Val() {
		share.Owner = ""
		share.FileId = ""
		share.FileName = ""
		share.Title = "来晚了,分享的文件已被删除"
		share.Size = 0
		share.SharingTime = ""
		return ""
	}
	shareInfo := cache.RedisClient.HGetAll(context.Background(), cache.ShareInfoKey(share.Uuid)).Val()
	share.Owner = shareInfo["Owner"]
	share.FileId = shareInfo["FileId"]
	share.FileName = shareInfo["FileName"]
	share.Title = shareInfo["Title"]
	share.Size, _ = strconv.ParseInt(shareInfo["Size"], 10, 64)
	share.SharingTime = shareInfo["SharingTime"]

	return shareInfo["downloadUrl"]
}

// CheckRedisExistsShare use title info to check, because title surely exsits
// when the share info store to redis
func (share *Share) CheckRedisExistsShare() bool {
	share.FileId, _ = cache.RedisClient.HGet(context.Background(), cache.ShareInfoKey(share.Uuid), "Title").Result()
	return share.FileId != "" || cache.RedisClient.SIsMember(context.Background(), cache.EmptyShare, share.Uuid).Val()
}

// DeleteShareInfoInRedis delete share info that in redis
func (share *Share) DeleteShareInfoInRedis() {
	_ = cache.RedisClient.ZRem(context.Background(), cache.DailyRankKey, share.Uuid)
	_ = cache.RedisClient.Del(context.Background(), cache.ShareInfoKey(share.Uuid)).Val()
}
