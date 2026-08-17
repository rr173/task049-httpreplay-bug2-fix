package replay

import "testing"

func TestProbe_ResponseBodyRoundTripsWithoutTruncation(t *testing.T) {
	r := New()
	req := Request{Method: "GET", Path: "/body"}
	if err := r.Record(req, Response{Status: 200, Body: []byte("abcde")}); err != nil {
		t.Fatal(err)
	}
	resp, err := r.Replay(req)
	if err != nil {
		t.Fatal(err)
	}
	if string(resp.Body) != "abcde" {
		t.Fatalf("body = %q, want %q", resp.Body, "abcde")
	}
}
