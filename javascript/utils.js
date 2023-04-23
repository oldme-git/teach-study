// 保存着实用的js函数

// 阿拉伯数字转中文
function arabToChinese(section) {
  let chnNumChar = ["零", "一", "二", "三", "四", "五", "六", "七", "八", "九"]
  let chnUnitChar = ["", "十", "百", "千", "万", "亿", "万亿", "亿亿"]
  let strIns = "", chnStr = ""
  let unitPos = 0
  let zero = true
  while (section > 0) {
    let v = section % 10
    if (v === 0) {
      if (!zero) {
        zero = true
        chnStr = chnNumChar[v] + chnStr
      }
    } else {
      zero = false
      strIns = chnNumChar[v]
      if (unitPos === 1 && v === 1 && chnStr !== "") {
        strIns = ""
      }
      strIns += chnUnitChar[unitPos]
      chnStr = strIns + chnStr
    }
    unitPos++
    section = Math.floor(section / 10)
  }
  return chnStr
}

// 取范围内的随机整数
function rndInt(min, max) {
  return Math.floor(Math.random() * (max - min + 1) ) + min
}

// 类似php的date函数
function formatDate(format, timestamp) {
  let a, jsDate = ((timestamp) ? new Date(timestamp * 1000) : new Date())
  let pad = function (n, c) {
    if ((n = n + "").length < c) {
      return new Array(++c - n.length).join("0") + n
    } else {
      return n
    }
  }
  let txtWeekdays = ["Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"]
  let txtOrdin = {1: "st", 2: "nd", 3: "rd", 21: "st", 22: "nd", 23: "rd", 31: "st"}
  let txtMonths = ["", "January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"]
  let f = {
    // Day
    d: function () {
      return pad(f.j(), 2)
    },
    D: function () {
      return f.l().substr(0, 3)
    },
    j: function () {
      return jsDate.getDate()
    },
    l: function () {
      return txtWeekdays[f.w()]
    },
    N: function () {
      return f.w() + 1
    },
    S: function () {
      return txtOrdin[f.j()] ? txtOrdin[f.j()] : 'th'
    },
    w: function () {
      return jsDate.getDay()
    },
    z: function () {
      return (jsDate - new Date(jsDate.getFullYear() + "/1/1")) / 864e5 >> 0
    },
    // Week
    W: function () {
      let a = f.z(), b = 364 + f.L() - a
      let nd2, nd = (new Date(jsDate.getFullYear() + "/1/1").getDay() || 7) - 1
      if (b <= 2 && ((jsDate.getDay() || 7) - 1) <= 2 - b) {
        return 1
      } else {
        if (a <= 2 && nd >= 4 && a >= (6 - nd)) {
          nd2 = new Date(jsDate.getFullYear() - 1 + "/12/31")
          return date("W", Math.round(nd2.getTime() / 1000))
        } else {
          return (1 + (nd <= 3 ? ((a + nd) / 7) : (a - (7 - nd)) / 7) >> 0)
        }
      }
    },

    // Month
    F: function () {
      return txtMonths[f.n()]
    },
    m: function () {
      return pad(f.n(), 2)
    },
    M: function () {
      return f.F().substr(0, 3)
    },
    n: function () {
      return jsDate.getMonth() + 1
    },
    t: function () {
      let n
      if ((n = jsDate.getMonth() + 1) == 2) {
        return 28 + f.L()
      } else {
        if (n & 1 && n < 8 || !(n & 1) && n > 7) {
          return 31
        } else {
          return 30
        }
      }
    },

    // Year
    L: function () {
      let y = f.Y()
      return (!(y & 3) && (y % 1e2 || !(y % 4e2))) ? 1 : 0
    },
    //o not supported yet
    Y: function () {
      return jsDate.getFullYear()
    },
    y: function () {
      return (jsDate.getFullYear() + "").slice(2)
    },

    // Time
    a: function () {
      return jsDate.getHours() > 11 ? "pm" : "am"
    },
    A: function () {
      return f.a().toUpperCase()
    },
    B: function () {
      // peter paul koch:
      let off = (jsDate.getTimezoneOffset() + 60) * 60
      let theSeconds = (jsDate.getHours() * 3600) + (jsDate.getMinutes() * 60) + jsDate.getSeconds() + off
      let beat = Math.floor(theSeconds / 86.4)
      if (beat > 1000) {
        beat -= 1000
      }
      if (beat < 0) {
        beat += 1000
      }
      if ((String(beat)).length == 1) {
        beat = "00" + beat
      }
      if ((String(beat)).length == 2) {
        beat = "0" + beat
      }
      return beat
    },
    g: function () {
      return jsDate.getHours() % 12 || 12
    },
    G: function () {
      return jsDate.getHours()
    },
    h: function () {
      return pad(f.g(), 2)
    },
    H: function () {
      return pad(jsDate.getHours(), 2)
    },
    i: function () {
      return pad(jsDate.getMinutes(), 2)
    },
    s: function () {
      return pad(jsDate.getSeconds(), 2)
    },
    //u not supported yet

    // Timezone
    //e not supported yet
    //I not supported yet
    O: function () {
      let t = pad(Math.abs(jsDate.getTimezoneOffset() / 60 * 100), 4)
      if (jsDate.getTimezoneOffset() > 0) {
        t = "-" + t
      } else {
        t = "+" + t
      }
      return t
    },
    P: function () {
      let O = f.O()
      return (O.substr(0, 3) + ":" + O.substr(3, 2))
    },
    //T not supported yet
    //Z not supported yet

    // Full Date/Time
    c: function () {
      return f.Y() + "-" + f.m() + "-" + f.d() + "T" + f.h() + ":" + f.i() + ":" + f.s() + f.P()
    },
    //r not supported yet
    U: function () {
      return Math.round(jsDate.getTime() / 1000)
    }
  }

  return format.replace(/([a-zA-Z])/g, function (t, s) {
    let ret
    if (t != s) {
      // escaped
      ret = s
    } else if (f[s]) {
      // a date function exists
      ret = f[s]()
    } else {
      // nothing special
      ret = s
    }
    return ret
  })
}
