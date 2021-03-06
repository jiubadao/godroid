package go_godroid

import (
	"code.google.com/p/go.mobile/bind/seq"
	"github.com/MarinX/godroid"
)

func proxy_ListeAndServe(out, in *seq.Buffer) {
	godroid.ListeAndServe()
}

func proxy_SayGo(out, in *seq.Buffer) {
	param_txt := in.ReadUTF16()
	godroid.SayGo(param_txt)
}

func init() {
	seq.Register("godroid", 1, proxy_ListeAndServe)
	seq.Register("godroid", 2, proxy_SayGo)
}

