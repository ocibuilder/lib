/*
Copyright 2019 BlackRock, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package dummy

import "github.com/ocibuilder/api/apis/beval/v1alpha1"

// Spec is a dummy spec
var Spec = v1alpha1.OCIBuilderSpec{
	Build: BuildSpec,
	Login: LoginSpec,
	Push:  PushSpec,
}

// BuildSpec is a dummy build spec
var BuildSpec = &v1alpha1.BuildSpec{
	Steps: []v1alpha1.BuildStep{
		{
			ImageMetadata: &v1alpha1.ImageMetadata{
				Name: "test-build",
			},
			Stages: []v1alpha1.Stage{
				{
					Base: v1alpha1.Base{
						Image: "alpine",
					},
					ImageMetadata: &v1alpha1.ImageMetadata{
						Name: "stage-one",
					},
					Cmd: []v1alpha1.BuildTemplateStep{
						{
							Docker: &v1alpha1.DockerStep{
								Inline: []string{"echo", "done"},
							},
						},
					},
				},
			},
			BuildContext: &v1alpha1.BuildContext{
				LocalContext: &v1alpha1.LocalContext{
					ContextPath: ".",
				},
			},
		},
	},
}

// LoginSpec is a dummy login spec
var LoginSpec = []v1alpha1.LoginSpec{
	{
		Registry: "example-registry",
		Token:    "ThiSiSalOgInToK3N",
		Creds: v1alpha1.RegistryCreds{
			Plain: v1alpha1.PlainCreds{
				Username: "username",
				Password: "password",
			},
		},
	},
}

// PushSpec is a dummy push spec
var PushSpec = []v1alpha1.PushSpec{
	{
		Registry: "example-registry",
		Image:    "example-image",
		Tag:      "1.0.0",
	},
}

// This key is NOT used anywhere other than this unit test
var TestPrivKey = `-----BEGIN PGP PRIVATE KEY BLOCK-----
Version: BCPG C# v1.6.1.0

lQOsBF4m0SQBCAC+VXNpTIglJlXIeEaseUL3aTqJmWnJE1Vpu7MYqT9rCKtnKlcN
BU2+WfcrG9ZJD/zUETxtw2m8nCZH/0K7XxjZPLo0qDfbM/giP9EcSJzmDeaAUMEo
buG/M5yxHReuysGZI/2X3Fw5swMr1mGOTSf6JLY6xecqlxgpI/N1IEWIKGSUmwlZ
fzhuBuV0EATov5zJ7XHdDUljrP3EdhJ3nurPwaUWkFjVeZEgqzj8QP0u8dIKQd4R
m/qWeR/cVa6btY8r2t4Cay4/ER8iYAHeqJHJlJmtCys+xEFzVwcemw27D+Mq4Ck6
wzc+3eypF2JJdrhicBqJ5JauG9NUCi8rynN7ABEBAAH/AwMCpW8FKB/oQuxgfeCG
X50b0fVmpRi+AESOGKDM/uUSISXQZUPxWojoKPG61E43oAD8utzpsI7TnnbjH1os
bEilVH/6QesYVXqMZIhMPdVyFPnTEr05MwkwaA4UeXLGHX5JsGd0l3nZtFGQlDp8
Tyxj7nUSUGEnQRmNZWzMnD6wfiMuaB2nfPAYjFPPeoVF09FGXpJcMPunbNenhinH
76I8M/OWFiUcBg6pEOy64ZoVG1sblKVcxC2Mv1g0koLQAANGJ4M6mjjQfJCcL0MW
Qd8C7bupd3m0Ph/S6LUPmH2ljmUhtaf42VNmMK9MAVcmxXYEyEefPCB3PKYBS7M1
hlzjwnpkscB0pzzfaq4AdveQMujHNG0rWIaKvvL4TzDauWnkYp0FnylFq8IewIxO
eTiflSjM/eWmAbGtUbonUMajokaTmR+DftJpeb0TBgFWKsj7bGf8xOHHDGl2T9fZ
zQOOdVO6ACNs60hj/hozv9sMFOTaJk1zX7kFXGJJiFeB+aCbyRrVKTSg4CKQFpXn
pAPDwc2KhEZZ1hAKJvL/mC565RJJvAveWcB0CB7w+0QOMkFrOTaqMVAwVvYop1xN
/qAzrzK3+5RZwb3ajbk7ShGE6JgdauYO2CgLPP0+jy36Iw+pgyTZYrLlH9U9k1ID
hsFYO8nhk+IkWc2RjIDs5ol3ym2DpHBjakWEaGwGiNfFWvxBH3Cp568J81CS6bSu
p7XLZLB97gBv8VAIP2t1UlJmnF/NKaBU76oacIccFRHfSIm+FE/crHYn+jyRZHnS
6y1tg5F4D1xkuKv1P3o1rxAmxkaifgHMCiQWGs3BsVFJe7OKXvIkWcuvnlFOevkc
dHUAs713JXWJYDQHDpKOo3c0iu1mseFCvdPjOton1bQAiQEcBBABAgAGBQJeJtEk
AAoJEE0wzYTUb94tRj8IALeVuBaZMb3HVRd589hxLYXNlzaK4WuMyOZavXzOLzji
bhPceiq3LRsXFY4U1xx9CtzyhTb8t0QrlLYgrTWXNvovezXlrDPuWH7J+5jPyy4o
3CAKGqTL+pVBRVM3MI+4D/wRatKM8uc08iCNJmuZI55sAmbZJR8IeQCgBzGf3cY/
0WxIKje8zQHHms+M3T3sQul07OoDD1qAVVWtWbbLPah36u18Gc77GiC2DtVoi8ux
m6LlB08sbpjUhjmxF+A34jPuKsLVP/gfyGktMQ4phtDL3T3cRstbNKDfX/IY9fAF
ZH7U4zGg6Wi5dl1oEMoLLavkiYK4Czwf9pRIpxhNIpE=
=fQz/
-----END PGP PRIVATE KEY BLOCK-----
`

var TestPubKey = `-----BEGIN PGP PUBLIC KEY BLOCK-----
Version: BCPG C# v1.6.1.0

mQENBF4m0SQBCAC+VXNpTIglJlXIeEaseUL3aTqJmWnJE1Vpu7MYqT9rCKtnKlcN
BU2+WfcrG9ZJD/zUETxtw2m8nCZH/0K7XxjZPLo0qDfbM/giP9EcSJzmDeaAUMEo
buG/M5yxHReuysGZI/2X3Fw5swMr1mGOTSf6JLY6xecqlxgpI/N1IEWIKGSUmwlZ
fzhuBuV0EATov5zJ7XHdDUljrP3EdhJ3nurPwaUWkFjVeZEgqzj8QP0u8dIKQd4R
m/qWeR/cVa6btY8r2t4Cay4/ER8iYAHeqJHJlJmtCys+xEFzVwcemw27D+Mq4Ck6
wzc+3eypF2JJdrhicBqJ5JauG9NUCi8rynN7ABEBAAG0AIkBHAQQAQIABgUCXibR
JAAKCRBNMM2E1G/eLUY/CAC3lbgWmTG9x1UXefPYcS2FzZc2iuFrjMjmWr18zi84
4m4T3Hoqty0bFxWOFNccfQrc8oU2/LdEK5S2IK01lzb6L3s15awz7lh+yfuYz8su
KNwgChqky/qVQUVTNzCPuA/8EWrSjPLnNPIgjSZrmSOebAJm2SUfCHkAoAcxn93G
P9FsSCo3vM0Bx5rPjN097ELpdOzqAw9agFVVrVm2yz2od+rtfBnO+xogtg7VaIvL
sZui5QdPLG6Y1IY5sRfgN+Iz7irC1T/4H8hpLTEOKYbQy9093EbLWzSg31/yGPXw
BWR+1OMxoOlouXZdaBDKCy2r5ImCuAs8H/aUSKcYTSKR
=oqhx
-----END PGP PUBLIC KEY BLOCK-----
`
