package main

type FrequencyTracker struct {
	mp, freq map[int]int
}

func Constructor() FrequencyTracker {
	return FrequencyTracker{map[int]int{}, map[int]int{}}
}

func (this *FrequencyTracker) Add(number int) {
	switch this.freq[this.mp[number]] {
	case 0:
	case 1:
		delete(this.freq, this.mp[number])
	default:
		this.freq[this.mp[number]]--
	}
	this.mp[number]++
	this.freq[this.mp[number]]++
}

func (this *FrequencyTracker) DeleteOne(number int) {
	if _, ok := this.mp[number]; ok {
		if this.freq[this.mp[number]] == 1 {
			delete(this.freq, this.mp[number])
		} else {
			this.freq[this.mp[number]]--
		}
		this.mp[number]--
		if this.mp[number] == 0 {
			delete(this.mp, number)
		} else {
			this.freq[this.mp[number]]++
		}
	}
}

func (this *FrequencyTracker) HasFrequency(frequency int) bool {
	_, ok := this.freq[frequency]
	return ok
}
