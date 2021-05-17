package TimeUnit

import "time"

type TimeUnit int64

const (
	NANOSECONDS  = time.Nanosecond
	MICROSECONDS = time.Microsecond
	MILLISECONDS = time.Millisecond
	SECONDS      = time.Second
	MINUTES      = time.Minute
	HOURS        = time.Hour
	DAYS         = 24 * HOURS
)

//convert(long sourceDuration, TimeUnit sourceUnit)
func Convert(target time.Duration, sourceDuration int64, sourceUnit time.Duration) int64 {
	rate := float64(sourceUnit) / float64(target)
	result := int64(rate * float64(sourceDuration))
	return result
}

//sleep(long timeout)
func Sleep(sourceUnit time.Duration, timeout int64) {
	duration := time.Duration(timeout)
	time.Sleep(sourceUnit * duration)
	//switch sourceUnit {
	//case NANOSECONDS:
	//case MICROSECONDS:
	//case MILLISECONDS:
	//case SECONDS:
	//case	MINUTES:
	//case	HOURS:
	//case	DAYS:
	//}
}

//timedJoin(Thread thread, long timeout)

//timedWait(Object obj, long timeout)
//toDays(long duration)
func ToDay(sourceUnit time.Duration, duration int64) int64 {
	result := Convert(DAYS, duration, sourceUnit)
	return result
}

//toHours(long duration)
func ToHours(sourceUnit time.Duration, duration int64) int64 {
	result := Convert(HOURS, duration, sourceUnit)
	return result
}

//toMicros(long duration)
func ToMicros(sourceUnit time.Duration, duration int64) int64 {
	result := Convert(MICROSECONDS, duration, sourceUnit)
	return result
}

//toMillis(long duration)
func ToMillis(sourceUnit time.Duration, duration int64) int64 {
	result := Convert(MILLISECONDS, duration, sourceUnit)
	return result
}

//toMinutes(long duration)
func ToMinutes(sourceUnit time.Duration, duration int64) int64 {
	result := Convert(MINUTES, duration, sourceUnit)
	return result
}

//toNanos(long duration)
func ToNanos(sourceUnit time.Duration, duration int64) int64 {
	result := Convert(NANOSECONDS, duration, sourceUnit)
	return result
}

//toSeconds(long duration)
func ToSeconds(sourceUnit time.Duration, duration int64) int64 {
	result := Convert(SECONDS, duration, sourceUnit)
	return result
}

//static TimeUnit	valueOf(String name)
//func ValueOf(name string)  TimeUnit{
//}
//static TimeUnit[]	values()
