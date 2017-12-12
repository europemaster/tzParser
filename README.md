# tzParser


## requirements

- params:    
  -t: time format (go time reference) 
    example: t := time.Time(*format*, generated regex (type str)) 
  -from: current tz (optional: auto detect tz if possible)
  -to: which tz to convert

- model:
  func getTS(logMessage)
    func getRegexFromFormat(format (type string))
      str := regex.MatchStr...
    func convertToTime(timeFormat)
  **returns** ts (type time.Time)
  
 - format example:
 **2017/06/09 17:03:15.998 ...log message...**
