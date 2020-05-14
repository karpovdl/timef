# timef

The package contains a set of methods for casting dates from one view to another.

## Sample

An example of working with the method **ToFormat**.

```
// 2020-05-20 00:00:00 -> 2020-05-20 00:00:00 
// Input date string
// > value: 2020-05-20 00:00:00
// > layout: 2006-01-02 15:04:05
// > format: YYYY-MM-DD HH24:MI:SS
outputDate, _ := timef.ToFormat("2020-05-20 00:00:00", timef.StampDashDateLongYearAtBegin, timef.GetFormat(FormatDateLongYearAtBegin21))
// Output date string
// > value: 2020-05-20 00:00:00
```

```
// 2020/05/20 00:00:00 -> 2020-05-20 00:00:00 
// Input date string
// > value: 2020/05/20 00:00:00
// > layout: 2006/01/02 15:04:05
// > format: YYYY-MM-DD HH24:MI:SS
outputDate, _ := timef.ToFormat("2020/05/20 00:00:00", timef.StampSlashDateLongYearAtBegin, timef.GetFormat(FormatDateLongYearAtBegin21))
// Output date string
// > value: 2020-05-20 00:00:00
```

```
// 2020.05.20 00:00:00 -> 2020-05-20 00:00:00 
// Input date string
// > value: 2020.05.20 00:00:00
// > layout: 2006.01.02 15:04:05
// > format: YYYY-MM-DD HH24:MI:SS
outputDate, _ := timef.ToFormat("2020.05.20 00:00:00", timef.StampDotDateLongYearAtBegin, timef.GetFormat(FormatDateLongYearAtBegin21))
// Output date string
// > value: 2020-05-20 00:00:00
```

```
// 20200520 00:00:00 -> 2020-05-20 00:00:00 
// Input date string
// > value: 20200520 00:00:00
// > layout: 20060102 15:04:05
// > format: YYYY-MM-DD HH24:MI:SS
outputDate, _ := timef.ToFormat("20200520 00:00:00", timef.StampDateLongYearAtBegin, timef.GetFormat(FormatDateLongYearAtBegin21))
// Output date string
// > value: 2020-05-20 00:00:00
```

```
// 20.05.2020 00:00:00 -> 2020-05-20 00:00:00 
// Input date string
// > value: 20.05.2020 00:00:00
// > layout: 02.01.2006 15:04:05
// > format: YYYY-MM-DD HH24:MI:SS
outputDate, _ := timef.ToFormat("20.05.2020 00:00:00", timef.StampDotDateLongYearAtEnd, timef.GetFormat(FormatDateLongYearAtBegin21))
// Output date string
// > value: 2020-05-20 00:00:00
```

```
// 20.05.20 00:00:00 -> 2020-05-20 00:00:00 
// Input date string
// > value: 20.05.20 00:00:00
// > layout: 02.01.06 15:04:05
// > format: YYYY-MM-DD HH24:MI:SS
outputDate, _ := timef.ToFormat("20.05.20 00:00:00", timef.StampDotDateYearAtEnd, timef.GetFormat(FormatDateLongYearAtBegin21))
// Output date string
// > value: 2020-05-20 00:00:00
```

```
// 2020-05-20 -> 2020-05-20 00:00:00 
// Input date string
// > value: 2020-05-20
// > layout: 2006-01-02
// > format: YYYY-MM-DD HH24:MI:SS
outputDate, _ := timef.ToFormat("2020-05-20", timef.StampDashDayLongYearAtBegin, timef.GetFormat(FormatDateLongYearAtBegin21))
// Output date string
// > value: 2020-05-20 00:00:00
```

```
// 20.05.2020 -> 2020-05-20 00:00:00 
// Input date string
// > value: 20.05.2020
// > layout: 02.01.2006
// > format: YYYY-MM-DD HH24:MI:SS
outputDate, _ := timef.ToFormat("20.05.2020", timef.StampDotDayLongYearAtEnd, timef.GetFormat(FormatDateLongYearAtBegin21))
// Output date string
// > value: 2020-05-20 00:00:00
```

```
// 20.05.20 -> 2020-05-20 00:00:00
// Input date string
// > value: 20.05.20
// > layout: 02.01.06
// > format: YYYY-MM-DD HH24:MI:SS
outputDate, _ := timef.ToFormat("20.05.20", timef.StampDotDayYearAtEnd, timef.GetFormat(FormatDateLongYearAtBegin21))
// Output date string
// > value: 2020-05-20 00:00:00
```