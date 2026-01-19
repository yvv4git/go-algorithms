package main

import "testing"

func TestCodec(t *testing.T) {
	obj := Constructor()
	longUrl := "https://leetcode.com/problems/design-tinyurl"

	// Test encode
	shortUrl := obj.encode(longUrl)
	if shortUrl == "" {
		t.Errorf("encode returned empty string")
	}
	if shortUrl == longUrl {
		t.Errorf("encode should return different URL")
	}

	// Test decode
	decoded := obj.decode(shortUrl)
	if decoded != longUrl {
		t.Errorf("decode failed: expected %s, got %s", longUrl, decoded)
	}

	// Test same URL encoded again returns same short
	shortUrl2 := obj.encode(longUrl)
	if shortUrl != shortUrl2 {
		t.Errorf("same URL should return same short URL: %s != %s", shortUrl, shortUrl2)
	}

	// Test different URL
	longUrl2 := "https://example.com"
	shortUrl3 := obj.encode(longUrl2)
	if shortUrl3 == shortUrl {
		t.Errorf("different URLs should have different short URLs")
	}
	decoded2 := obj.decode(shortUrl3)
	if decoded2 != longUrl2 {
		t.Errorf("decode failed for second URL: expected %s, got %s", longUrl2, decoded2)
	}
}
