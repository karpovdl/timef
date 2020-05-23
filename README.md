# timef

[![License][1]][2] [![golang][10]][11] [![codebeat][20]][21] [![CodeFactor][22]][23] 

[1]: https://img.shields.io/badge/license-MIT-blue.svg?label=License&maxAge=86400 "License"
[2]: ./LICENSE

[10]: https://img.shields.io/badge/golang-1.14.2-blue.svg?style=flat "Golang"
[11]: https://golang.org

[20]: https://codebeat.co/badges/a08a9126-60d7-4cee-b266-8277a7885467 "CODEBEAT"
[21]: https://codebeat.co/projects/github-com-karpovdl-timef-master

[22]: https://www.codefactor.io/repository/github/karpovdl/timef/badge "CodeFactor"
[23]: https://www.codefactor.io/repository/github/karpovdl/timef

[![](resource/timef.png)](https://github.com/karpovdl/timef)

The package contains a set of methods for casting dates from one view to another.

## Sample

An example of working with the method **ToFormat**.

```
// 2020-05-20 00:00:00 -> 2020-05-20 00:00:00 
// Input date string
// > value: 2020-05-20 00:00:00
// > layout: 2006-01-02 15:04:05
// > format: YYYY-MM-DD HH24:MI:SS
outputDate, _ := timef.ToFormat("2020-05-20 00:00:00", timef.StampDashDateLongYearAtBegin, timef.Format[FormatDateLongYearAtBegin21])
// Output date string
// > value: 2020-05-20 00:00:00
```

```
// 2020/05/20 00:00:00 -> 2020-05-20 00:00:00 
// Input date string
// > value: 2020/05/20 00:00:00
// > layout: 2006/01/02 15:04:05
// > format: YYYY-MM-DD HH24:MI:SS
outputDate, _ := timef.ToFormat("2020/05/20 00:00:00", timef.StampSlashDateLongYearAtBegin, timef.Format[FormatDateLongYearAtBegin21])
// Output date string
// > value: 2020-05-20 00:00:00
```

```
// 2020.05.20 00:00:00 -> 2020-05-20 00:00:00 
// Input date string
// > value: 2020.05.20 00:00:00
// > layout: 2006.01.02 15:04:05
// > format: YYYY-MM-DD HH24:MI:SS
outputDate, _ := timef.ToFormat("2020.05.20 00:00:00", timef.StampDotDateLongYearAtBegin, timef.Format[FormatDateLongYearAtBegin21])
// Output date string
// > value: 2020-05-20 00:00:00
```

```
// 20200520 00:00:00 -> 2020-05-20 00:00:00 
// Input date string
// > value: 20200520 00:00:00
// > layout: 20060102 15:04:05
// > format: YYYY-MM-DD HH24:MI:SS
outputDate, _ := timef.ToFormat("20200520 00:00:00", timef.StampDateLongYearAtBegin, timef.Format[FormatDateLongYearAtBegin21])
// Output date string
// > value: 2020-05-20 00:00:00
```

```
// 20.05.2020 00:00:00 -> 2020-05-20 00:00:00 
// Input date string
// > value: 20.05.2020 00:00:00
// > layout: 02.01.2006 15:04:05
// > format: YYYY-MM-DD HH24:MI:SS
outputDate, _ := timef.ToFormat("20.05.2020 00:00:00", timef.StampDotDateLongYearAtEnd, timef.Format[FormatDateLongYearAtBegin21])
// Output date string
// > value: 2020-05-20 00:00:00
```

```
// 20.05.20 00:00:00 -> 2020-05-20 00:00:00 
// Input date string
// > value: 20.05.20 00:00:00
// > layout: 02.01.06 15:04:05
// > format: YYYY-MM-DD HH24:MI:SS
outputDate, _ := timef.ToFormat("20.05.20 00:00:00", timef.StampDotDateYearAtEnd, timef.Format[FormatDateLongYearAtBegin21])
// Output date string
// > value: 2020-05-20 00:00:00
```

```
// 2020-05-20 -> 2020-05-20 00:00:00 
// Input date string
// > value: 2020-05-20
// > layout: 2006-01-02
// > format: YYYY-MM-DD HH24:MI:SS
outputDate, _ := timef.ToFormat("2020-05-20", timef.StampDashDayLongYearAtBegin, timef.Format[FormatDateLongYearAtBegin21])
// Output date string
// > value: 2020-05-20 00:00:00
```

```
// 20.05.2020 -> 2020-05-20 00:00:00 
// Input date string
// > value: 20.05.2020
// > layout: 02.01.2006
// > format: YYYY-MM-DD HH24:MI:SS
outputDate, _ := timef.ToFormat("20.05.2020", timef.StampDotDayLongYearAtEnd, timef.Format[FormatDateLongYearAtBegin21])
// Output date string
// > value: 2020-05-20 00:00:00
```

```
// 20.05.20 -> 2020-05-20 00:00:00
// Input date string
// > value: 20.05.20
// > layout: 02.01.06
// > format: YYYY-MM-DD HH24:MI:SS
outputDate, _ := timef.ToFormat("20.05.20", timef.StampDotDayYearAtEnd, timef.Format[FormatDateLongYearAtBegin21])
// Output date string
// > value: 2020-05-20 00:00:00
```
