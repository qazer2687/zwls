## zwls
a zero-width link shortener

---

### about
Hashes URLs with FNV-1 32-bit and the integer is encoded into base-4 with zero-width unicode characters as digits. The resulting slug is invisible when rendered but remains a valid URL.

### endpoints

`POST /shorten` - body is URL, returns zero-width slug  
`GET /{slug}` - 301 redirect to original URL  
`GET /` - health check