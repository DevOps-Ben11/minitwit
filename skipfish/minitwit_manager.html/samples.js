var mime_samples = [
  { 'mime': 'application/javascript', 'samples': [
    { 'url': 'http://209.38.250.33/assets/index-DOMSawVM.js', 'dir': '_m0/0', 'linked': 2, 'len': 246451 } ]
  },
  { 'mime': 'image/x-ms-bmp', 'samples': [
    { 'url': 'http://209.38.250.33/icon.ico', 'dir': '_m1/0', 'linked': 2, 'len': 1406 } ]
  },
  { 'mime': 'text/css', 'samples': [
    { 'url': 'http://209.38.250.33/assets/index-CXRwShXL.css', 'dir': '_m2/0', 'linked': 2, 'len': 3377 } ]
  },
  { 'mime': 'text/html', 'samples': [
    { 'url': 'http://209.38.250.33/', 'dir': '_m3/0', 'linked': 2, 'len': 451 },
    { 'url': 'http://209.38.250.33/assets/', 'dir': '_m3/1', 'linked': 2, 'len': 115 } ]
  },
  { 'mime': 'text/plain', 'samples': [
    { 'url': 'http://209.38.250.33/', 'dir': '_m4/0', 'linked': 2, 'len': 15 },
    { 'url': 'http://209.38.250.33/assets/sfi9876', 'dir': '_m4/1', 'linked': 2, 'len': 19 } ]
  }
];

var issue_samples = [
  { 'severity': 1, 'type': 20101, 'samples': [
    { 'url': 'http://209.38.250.33/`echo%24{IFS}skip12``echo%24{IFS}34fish`', 'extra': 'Shell injection (spec)', 'sid': '0', 'dir': '_i0/0' },
    { 'url': 'http://209.38.250.33/assets/?_test1=c:\x5cwindows\x5csystem32\x5ccmd.exe&_test2=/etc/passwd&_test3=|/bin/sh&_test4=(SELECT%20*%20FROM%20nonexistent)%20--&_test5=\x3e/no/such/file&_test6=\x3cscript\x3ealert(1)\x3c/script\x3e&_test7=javascript:alert(1)', 'extra': 'IPS check', 'sid': '0', 'dir': '_i0/1' },
    { 'url': 'http://209.38.250.33/assets/index-DOMSawVM.js/`sleep%24{IFS}5`', 'extra': 'Shell injection (spec)', 'sid': '0', 'dir': '_i0/2' },
    { 'url': 'http://209.38.250.33/icon.ico/`sleep%24{IFS}5`', 'extra': 'Shell injection (spec)', 'sid': '0', 'dir': '_i0/3' } ]
  },
  { 'severity': 0, 'type': 10802, 'samples': [
    { 'url': 'http://209.38.250.33/', 'extra': 'text/plain', 'sid': '0', 'dir': '_i1/0' },
    { 'url': 'http://209.38.250.33/assets/sfi9876', 'extra': 'text/plain', 'sid': '0', 'dir': '_i1/1' },
    { 'url': 'http://209.38.250.33/assets/index-CXRwShXL.css', 'extra': 'text/plain', 'sid': '0', 'dir': '_i1/2' },
    { 'url': 'http://209.38.250.33/assets/index-DOMSawVM.js', 'extra': 'text/plain', 'sid': '0', 'dir': '_i1/3' },
    { 'url': 'http://209.38.250.33/icon.ico', 'extra': 'text/plain', 'sid': '0', 'dir': '_i1/4' } ]
  },
  { 'severity': 0, 'type': 10801, 'samples': [
    { 'url': 'http://209.38.250.33/icon.ico', 'extra': 'image/x-ms-bmp', 'sid': '0', 'dir': '_i2/0' } ]
  },
  { 'severity': 0, 'type': 10205, 'samples': [
    { 'url': 'http://209.38.250.33/sfi9876', 'extra': '', 'sid': '0', 'dir': '_i3/0' },
    { 'url': 'http://209.38.250.33/assets/sfi9876', 'extra': '', 'sid': '0', 'dir': '_i3/1' } ]
  }
];

