package main

import "container/heap"

// Twitter 推特
type Twitter struct {
	users map[int]*User
	time  int
}

// User 用户
type User struct {
	id        int
	followee  map[int]*User
	tweetList *Tweet
}

// Tweet 推文
type Tweet struct {
	tweetId int
	time    int
	next    *Tweet
}

func Constructor() Twitter {
	return Twitter{users: make(map[int]*User)}
}

// getTime 获取当前系统时间戳，并对时间戳 +1
func (t *Twitter) getTime() int {
	t.time++
	return t.time
}

// getUser 获取指定用户 ID 的 User 对象
func (t *Twitter) getUser(id int) *User {
	if u, ok := t.users[id]; ok {
		return u
	}
	t.users[id] = &User{id: id, followee: make(map[int]*User)}
	return t.users[id]
}

// PostTweet 根据给定的 tweetId 和 userId 创建一条新推文。每次调用此函数都会使用一个不同的 tweetId
func (t *Twitter) PostTweet(userId, tweetId int) {
	tweet := Tweet{tweetId: tweetId, time: t.getTime()}
	user := t.getUser(userId)

	// 在链表头插入新节点
	tweet.next = user.tweetList
	user.tweetList = &tweet
}

const TimeLineSize = 10 // 信息流包括用户的关注人（包括自己）的最近 10 条推文

// GetNewsFeed 检索当前用户新闻推送中最近 10 条推文的 ID 。新闻推送中的每一项都必须是由用户关注的人或者是用户自己发布的推文。推文按照时间顺序由最近到最远排序
func (t *Twitter) GetNewsFeed(userId int) []int {
	user := t.getUser(userId)

	tweets := make([]*Tweet, 0, len(user.followee)+1)
	if user.tweetList != nil { // 信息流需包含用户自己发布的推文
		tweets = append(tweets, user.tweetList)
	}
	for _, followee := range user.followee {
		if followee.tweetList != nil {
			tweets = append(tweets, followee.tweetList)
		}
	}

	res := make([]int, 0, TimeLineSize)
	tweetHeap := TweetHeap(tweets)
	h := &tweetHeap
	heap.Init(h)
	for len(res) < 10 && h.Len() > 0 {
		text := heap.Pop(h).(*Tweet)
		res = append(res, text.tweetId)
		if text.next != nil {
			heap.Push(h, text.next)
		}
	}
	return res
}

// Follow ID 为 followerId 的用户开始关注 ID 为 followeeId 的用户
func (t *Twitter) Follow(followerId, followeeId int) {
	if followerId == followeeId { // 用户无法关注自己
		return
	}
	follower, followee := t.getUser(followerId), t.getUser(followeeId)
	follower.followee[followeeId] = followee
}

// Unfollow ID 为 followerId 的用户不再关注 ID 为 followeeId 的用户
func (t *Twitter) Unfollow(followerId, followeeId int) {
	if followerId == followeeId { // 用户无法取关自己
		return
	}
	follower := t.getUser(followerId)
	delete(follower.followee, followeeId)
}

type TweetHeap []*Tweet

func (h TweetHeap) Len() int           { return len(h) }
func (h TweetHeap) Less(i, j int) bool { return h[i].time > h[j].time }
func (h TweetHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *TweetHeap) Push(x any)        { *h = append(*h, x.(*Tweet)) }
func (h *TweetHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
