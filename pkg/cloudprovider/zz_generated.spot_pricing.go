//go:build !ignore_autogenerated

/*
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

package cloudprovider

import "time"

// generated at 2023-01-21T22:31:12Z for us-east-1

var initialSpotPriceUpdate, _ = time.Parse(time.RFC3339, "2023-01-21T22:31:12Z")
var initialSpotPrices = map[string]float64{
	// a1 family
	"a1.2xlarge": 0.077900, "a1.4xlarge": 0.153900, "a1.large": 0.039900, "a1.medium": 0.009400,
	"a1.metal": 0.134300, "a1.xlarge": 0.048800,
	// c1 family
	"c1.medium": 0.035800, "c1.xlarge": 0.520000,
	// c3 family
	"c3.2xlarge": 0.155300, "c3.4xlarge": 0.312200, "c3.8xlarge": 0.586200, "c3.large": 0.038000,
	"c3.xlarge": 0.104600,
	// c4 family
	"c4.2xlarge": 0.193600, "c4.4xlarge": 0.293800, "c4.8xlarge": 0.595200, "c4.large": 0.035600,
	"c4.xlarge": 0.104700,
	// c5 family
	"c5.12xlarge": 0.874500, "c5.18xlarge": 1.510200, "c5.24xlarge": 1.554600, "c5.2xlarge": 0.189800,
	"c5.4xlarge": 0.269100, "c5.9xlarge": 0.687400, "c5.large": 0.039400, "c5.metal": 1.560400,
	"c5.xlarge": 0.086400,
	// c5a family
	"c5a.12xlarge": 0.795100, "c5a.16xlarge": 1.036400, "c5a.24xlarge": 1.564300, "c5a.2xlarge": 0.174900,
	"c5a.4xlarge": 0.335900, "c5a.8xlarge": 0.531400, "c5a.large": 0.037700, "c5a.xlarge": 0.085400,
	// c5ad family
	"c5ad.12xlarge": 0.789100, "c5ad.16xlarge": 1.041800, "c5ad.24xlarge": 1.679400, "c5ad.2xlarge": 0.182200,
	"c5ad.4xlarge": 0.336300, "c5ad.8xlarge": 0.670500, "c5ad.large": 0.032400, "c5ad.xlarge": 0.089200,
	// c5d family
	"c5d.12xlarge": 1.103300, "c5d.18xlarge": 1.394700, "c5d.24xlarge": 1.573100, "c5d.2xlarge": 0.150000,
	"c5d.4xlarge": 0.341400, "c5d.9xlarge": 0.767400, "c5d.large": 0.032500, "c5d.metal": 1.562700,
	"c5d.xlarge": 0.095500,
	// c5n family
	"c5n.18xlarge": 1.259800, "c5n.2xlarge": 0.195000, "c5n.4xlarge": 0.374100, "c5n.9xlarge": 0.692100,
	"c5n.large": 0.035200, "c5n.metal": 1.170700, "c5n.xlarge": 0.084000,
	// c6a family
	"c6a.12xlarge": 1.109200, "c6a.16xlarge": 1.230400, "c6a.24xlarge": 1.632300, "c6a.2xlarge": 0.174200,
	"c6a.32xlarge": 2.176400, "c6a.48xlarge": 3.285000, "c6a.4xlarge": 0.334200, "c6a.8xlarge": 0.603700,
	"c6a.large": 0.034000, "c6a.metal": 5.225200, "c6a.xlarge": 0.085500,
	// c6g family
	"c6g.12xlarge": 0.816200, "c6g.16xlarge": 1.088200, "c6g.2xlarge": 0.164300, "c6g.4xlarge": 0.272100,
	"c6g.8xlarge": 0.544100, "c6g.large": 0.039400, "c6g.medium": 0.028300, "c6g.metal": 1.088200,
	"c6g.xlarge": 0.070400,
	// c6gd family
	"c6gd.12xlarge": 0.818400, "c6gd.16xlarge": 1.088200, "c6gd.2xlarge": 0.139600, "c6gd.4xlarge": 0.281700,
	"c6gd.8xlarge": 0.544100, "c6gd.large": 0.034200, "c6gd.medium": 0.017300, "c6gd.metal": 1.088200,
	"c6gd.xlarge": 0.068900,
	// c6gn family
	"c6gn.12xlarge": 0.951500, "c6gn.16xlarge": 1.088200, "c6gn.2xlarge": 0.173500, "c6gn.4xlarge": 0.278800,
	"c6gn.8xlarge": 0.726000, "c6gn.large": 0.034000, "c6gn.medium": 0.017000, "c6gn.xlarge": 0.075900,
	// c6i family
	"c6i.12xlarge": 0.845200, "c6i.16xlarge": 1.109400, "c6i.24xlarge": 1.659000, "c6i.2xlarge": 0.167000,
	"c6i.32xlarge": 2.192500, "c6i.4xlarge": 0.278600, "c6i.8xlarge": 0.556500, "c6i.large": 0.034200,
	"c6i.metal": 2.178900, "c6i.xlarge": 0.085000,
	// c6id family
	"c6id.12xlarge": 0.816200, "c6id.16xlarge": 1.088200, "c6id.24xlarge": 2.008600, "c6id.2xlarge": 0.169200,
	"c6id.32xlarge": 6.451200, "c6id.4xlarge": 0.273700, "c6id.8xlarge": 0.544100, "c6id.large": 0.045400,
	"c6id.metal": 2.176400, "c6id.xlarge": 0.072300,
	// c7g family
	"c7g.12xlarge": 0.857000, "c7g.16xlarge": 1.142600, "c7g.2xlarge": 0.142800, "c7g.4xlarge": 0.302700,
	"c7g.8xlarge": 0.684100, "c7g.large": 0.035700, "c7g.medium": 0.017900, "c7g.xlarge": 0.071400,
	// d2 family
	"d2.2xlarge": 0.414000, "d2.4xlarge": 0.828000, "d2.8xlarge": 1.656000, "d2.xlarge": 0.207000,
	// d3 family
	"d3.2xlarge": 0.299700, "d3.4xlarge": 0.599400, "d3.8xlarge": 1.198700, "d3.xlarge": 0.149700,
	// d3en family
	"d3en.12xlarge": 1.892600, "d3en.2xlarge": 0.315300, "d3en.4xlarge": 0.630900, "d3en.6xlarge": 0.997200,
	"d3en.8xlarge": 1.261700, "d3en.xlarge": 0.157800,
	// dl1 family
	"dl1.24xlarge": 3.932700,
	// f1 family
	"f1.16xlarge": 3.960000, "f1.2xlarge": 0.495000, "f1.4xlarge": 1.034700,
	// g2 family
	"g2.2xlarge": 0.650000, "g2.8xlarge": 2.600000,
	// g3 family
	"g3.16xlarge": 1.368000, "g3.4xlarge": 0.342000, "g3.8xlarge": 0.684000,
	// g3s family
	"g3s.xlarge": 0.285600,
	// g4ad family
	"g4ad.16xlarge": 1.121100, "g4ad.2xlarge": 0.162400, "g4ad.4xlarge": 0.350700, "g4ad.8xlarge": 0.556900,
	"g4ad.xlarge": 0.113600,
	// g4dn family
	"g4dn.12xlarge": 1.247800, "g4dn.16xlarge": 1.305600, "g4dn.2xlarge": 0.225600, "g4dn.4xlarge": 0.361200,
	"g4dn.8xlarge": 0.652800, "g4dn.metal": 2.518500, "g4dn.xlarge": 0.157800,
	// g5 family
	"g5.12xlarge": 1.701600, "g5.16xlarge": 1.228800, "g5.24xlarge": 2.635700, "g5.2xlarge": 0.363600,
	"g5.48xlarge": 5.257500, "g5.4xlarge": 0.487200, "g5.8xlarge": 0.846300, "g5.xlarge": 0.301800,
	// g5g family
	"g5g.16xlarge": 0.823200, "g5g.2xlarge": 0.219700, "g5g.4xlarge": 0.248400, "g5g.8xlarge": 0.511500,
	"g5g.metal": 0.823200, "g5g.xlarge": 0.126000,
	// h1 family
	"h1.16xlarge": 1.123200, "h1.2xlarge": 0.178200, "h1.4xlarge": 0.414000, "h1.8xlarge": 0.628900,
	// i2 family
	"i2.2xlarge": 0.511500, "i2.4xlarge": 1.023000, "i2.8xlarge": 2.046000, "i2.xlarge": 0.255900,
	// i3 family
	"i3.16xlarge": 1.674400, "i3.2xlarge": 0.257800, "i3.4xlarge": 0.590200, "i3.8xlarge": 0.878300,
	"i3.large": 0.046800, "i3.metal": 1.497600, "i3.xlarge": 0.136900,
	// i3en family
	"i3en.12xlarge": 1.627200, "i3en.24xlarge": 3.254400, "i3en.2xlarge": 0.328600, "i3en.3xlarge": 0.470600,
	"i3en.6xlarge": 1.015700, "i3en.large": 0.067800, "i3en.metal": 3.254400, "i3en.xlarge": 0.152000,
	// i4i family
	"i4i.16xlarge": 1.647300, "i4i.2xlarge": 0.205800, "i4i.32xlarge": 3.294700, "i4i.4xlarge": 0.430700,
	"i4i.8xlarge": 0.823800, "i4i.large": 0.051600, "i4i.metal": 3.294600, "i4i.xlarge": 0.102900,
	// im4gn family
	"im4gn.16xlarge": 1.746200, "im4gn.2xlarge": 0.218300, "im4gn.4xlarge": 0.436600, "im4gn.8xlarge": 0.873100,
	"im4gn.large": 0.054600, "im4gn.xlarge": 0.109100,
	// inf1 family
	"inf1.24xlarge": 1.426700, "inf1.2xlarge": 0.117800, "inf1.6xlarge": 0.355600, "inf1.xlarge": 0.074900,
	// is4gen family
	"is4gen.2xlarge": 0.345800, "is4gen.4xlarge": 0.691600, "is4gen.8xlarge": 1.383100,
	"is4gen.large": 0.086400, "is4gen.medium": 0.043200, "is4gen.xlarge": 0.196900,
	// m1 family
	"m1.large": 0.025900, "m1.medium": 0.046800, "m1.small": 0.020400, "m1.xlarge": 0.054900,
	// m2 family
	"m2.2xlarge": 0.411200, "m2.4xlarge": 0.127800, "m2.xlarge": 0.129600,
	// m3 family
	"m3.2xlarge": 0.160700, "m3.large": 0.031500, "m3.medium": 0.006700, "m3.xlarge": 0.062000,
	// m4 family
	"m4.10xlarge": 0.863700, "m4.16xlarge": 1.057300, "m4.2xlarge": 0.200800, "m4.4xlarge": 0.377200,
	"m4.large": 0.043400, "m4.xlarge": 0.070300,
	// m5 family
	"m5.12xlarge": 0.919600, "m5.16xlarge": 1.531900, "m5.24xlarge": 1.908600, "m5.2xlarge": 0.150300,
	"m5.4xlarge": 0.346000, "m5.8xlarge": 0.599000, "m5.large": 0.043000, "m5.metal": 1.632300,
	"m5.xlarge": 0.068400,
	// m5a family
	"m5a.12xlarge": 0.864800, "m5a.16xlarge": 1.091000, "m5a.24xlarge": 1.632300, "m5a.2xlarge": 0.200700,
	"m5a.4xlarge": 0.316300, "m5a.8xlarge": 0.672200, "m5a.large": 0.046700, "m5a.xlarge": 0.088400,
	// m5ad family
	"m5ad.12xlarge": 1.099200, "m5ad.16xlarge": 1.134900, "m5ad.24xlarge": 1.716900, "m5ad.2xlarge": 0.293000,
	"m5ad.4xlarge": 0.615600, "m5ad.8xlarge": 0.703300, "m5ad.large": 0.046900, "m5ad.xlarge": 0.130100,
	// m5d family
	"m5d.12xlarge": 1.092100, "m5d.16xlarge": 1.393700, "m5d.24xlarge": 1.924400, "m5d.2xlarge": 0.250500,
	"m5d.4xlarge": 0.348300, "m5d.8xlarge": 0.586000, "m5d.large": 0.037200, "m5d.metal": 1.632300,
	"m5d.xlarge": 0.079900,
	// m5dn family
	"m5dn.12xlarge": 1.091100, "m5dn.16xlarge": 1.873900, "m5dn.24xlarge": 1.632300, "m5dn.2xlarge": 0.319800,
	"m5dn.4xlarge": 0.658700, "m5dn.8xlarge": 0.544100, "m5dn.large": 0.038700, "m5dn.metal": 1.643400,
	"m5dn.xlarge": 0.185300,
	// m5n family
	"m5n.12xlarge": 1.977800, "m5n.16xlarge": 2.104700, "m5n.24xlarge": 2.012700, "m5n.2xlarge": 0.273800,
	"m5n.4xlarge": 0.549200, "m5n.8xlarge": 0.779700, "m5n.large": 0.055700, "m5n.metal": 1.632300,
	"m5n.xlarge": 0.068100,
	// m5zn family
	"m5zn.12xlarge": 0.816200, "m5zn.2xlarge": 0.161500, "m5zn.3xlarge": 0.399000, "m5zn.6xlarge": 0.410000,
	"m5zn.large": 0.087400, "m5zn.metal": 0.852600, "m5zn.xlarge": 0.234400,
	// m6a family
	"m6a.12xlarge": 1.604700, "m6a.16xlarge": 1.326500, "m6a.24xlarge": 1.721900, "m6a.2xlarge": 0.154700,
	"m6a.32xlarge": 2.285200, "m6a.48xlarge": 3.427900, "m6a.4xlarge": 0.380100, "m6a.8xlarge": 0.631700,
	"m6a.large": 0.035800, "m6a.metal": 3.427900, "m6a.xlarge": 0.085300,
	// m6g family
	"m6g.12xlarge": 0.857000, "m6g.16xlarge": 1.142600, "m6g.2xlarge": 0.145300, "m6g.4xlarge": 0.290000,
	"m6g.8xlarge": 0.604800, "m6g.large": 0.036500, "m6g.medium": 0.017900, "m6g.metal": 1.142600,
	"m6g.xlarge": 0.071400,
	// m6gd family
	"m6gd.12xlarge": 1.160500, "m6gd.16xlarge": 1.142600, "m6gd.2xlarge": 0.155100, "m6gd.4xlarge": 0.311300,
	"m6gd.8xlarge": 0.575700, "m6gd.large": 0.038400, "m6gd.medium": 0.017900, "m6gd.metal": 1.142600,
	"m6gd.xlarge": 0.085700,
	// m6i family
	"m6i.12xlarge": 1.048400, "m6i.16xlarge": 1.343100, "m6i.24xlarge": 1.794400, "m6i.2xlarge": 0.189400,
	"m6i.32xlarge": 2.289300, "m6i.4xlarge": 0.358000, "m6i.8xlarge": 0.608800, "m6i.large": 0.035800,
	"m6i.metal": 2.285200, "m6i.xlarge": 0.074000,
	// m6id family
	"m6id.12xlarge": 0.857000, "m6id.16xlarge": 1.162600, "m6id.24xlarge": 5.695200, "m6id.2xlarge": 0.187700,
	"m6id.32xlarge": 7.593600, "m6id.4xlarge": 0.358500, "m6id.8xlarge": 0.650600, "m6id.large": 0.035700,
	"m6id.metal": 2.285200, "m6id.xlarge": 0.072000,
	// p2 family
	"p2.16xlarge": 4.320000, "p2.8xlarge": 2.160000, "p2.xlarge": 0.397600,
	// p3 family
	"p3.16xlarge": 7.344000, "p3.2xlarge": 1.220000, "p3.8xlarge": 3.672000,
	// p4d family
	"p4d.24xlarge": 9.831800,
	// r3 family
	"r3.2xlarge": 0.155800, "r3.4xlarge": 0.316200, "r3.8xlarge": 0.522500, "r3.large": 0.040300,
	"r3.xlarge": 0.076600,
	// r4 family
	"r4.16xlarge": 1.088700, "r4.2xlarge": 0.200400, "r4.4xlarge": 0.292900, "r4.8xlarge": 0.546700,
	"r4.large": 0.043200, "r4.xlarge": 0.088800,
	// r5 family
	"r5.12xlarge": 0.862200, "r5.16xlarge": 1.561000, "r5.24xlarge": 2.237700, "r5.2xlarge": 0.167100,
	"r5.4xlarge": 0.346400, "r5.8xlarge": 0.722200, "r5.large": 0.037900, "r5.metal": 1.710000,
	"r5.xlarge": 0.071500,
	// r5a family
	"r5a.12xlarge": 0.997000, "r5a.16xlarge": 1.146600, "r5a.24xlarge": 1.710000, "r5a.2xlarge": 0.164200,
	"r5a.4xlarge": 0.292000, "r5a.8xlarge": 0.674900, "r5a.large": 0.042100, "r5a.xlarge": 0.114700,
	// r5ad family
	"r5ad.12xlarge": 1.072700, "r5ad.16xlarge": 1.597100, "r5ad.24xlarge": 1.938100, "r5ad.2xlarge": 0.169400,
	"r5ad.4xlarge": 0.443900, "r5ad.8xlarge": 0.631000, "r5ad.large": 0.036400, "r5ad.xlarge": 0.083700,
	// r5b family
	"r5b.12xlarge": 0.904400, "r5b.16xlarge": 1.237000, "r5b.24xlarge": 3.229900, "r5b.2xlarge": 0.257500,
	"r5b.4xlarge": 0.839600, "r5b.8xlarge": 0.663500, "r5b.large": 0.036700, "r5b.metal": 3.156500,
	"r5b.xlarge": 0.119600,
	// r5d family
	"r5d.12xlarge": 1.546600, "r5d.16xlarge": 1.645900, "r5d.24xlarge": 2.636700, "r5d.2xlarge": 0.199700,
	"r5d.4xlarge": 0.479700, "r5d.8xlarge": 0.861700, "r5d.large": 0.036600, "r5d.metal": 1.710000,
	"r5d.xlarge": 0.089600,
	// r5dn family
	"r5dn.12xlarge": 1.160400, "r5dn.16xlarge": 1.579700, "r5dn.24xlarge": 2.813100, "r5dn.2xlarge": 0.191500,
	"r5dn.4xlarge": 0.397200, "r5dn.8xlarge": 0.982000, "r5dn.large": 0.036700, "r5dn.metal": 3.662100,
	"r5dn.xlarge": 0.082600,
	// r5n family
	"r5n.12xlarge": 1.579100, "r5n.16xlarge": 2.363800, "r5n.24xlarge": 2.409600, "r5n.2xlarge": 0.192100,
	"r5n.4xlarge": 0.535800, "r5n.8xlarge": 0.688300, "r5n.large": 0.037400, "r5n.metal": 1.881200,
	"r5n.xlarge": 0.092100,
	// r6a family
	"r6a.12xlarge": 1.486400, "r6a.16xlarge": 1.301600, "r6a.24xlarge": 1.937500, "r6a.2xlarge": 0.149600,
	"r6a.32xlarge": 2.394100, "r6a.48xlarge": 3.996300, "r6a.4xlarge": 0.370500, "r6a.8xlarge": 0.602200,
	"r6a.large": 0.037400, "r6a.metal": 3.591100, "r6a.xlarge": 0.096200,
	// r6g family
	"r6g.12xlarge": 0.897800, "r6g.16xlarge": 1.476000, "r6g.2xlarge": 0.151600, "r6g.4xlarge": 0.339400,
	"r6g.8xlarge": 0.598500, "r6g.large": 0.045200, "r6g.medium": 0.018800, "r6g.metal": 1.197000,
	"r6g.xlarge": 0.074800,
	// r6gd family
	"r6gd.12xlarge": 0.897800, "r6gd.16xlarge": 1.197000, "r6gd.2xlarge": 0.163500, "r6gd.4xlarge": 0.368000,
	"r6gd.8xlarge": 0.598500, "r6gd.large": 0.037400, "r6gd.medium": 0.018700, "r6gd.metal": 1.197000,
	"r6gd.xlarge": 0.083200,
	// r6i family
	"r6i.12xlarge": 2.035700, "r6i.16xlarge": 1.341300, "r6i.24xlarge": 2.014400, "r6i.2xlarge": 0.160100,
	"r6i.32xlarge": 2.448400, "r6i.4xlarge": 0.412500, "r6i.8xlarge": 0.646100, "r6i.large": 0.037400,
	"r6i.metal": 2.394100, "r6i.xlarge": 0.075300,
	// r6id family
	"r6id.12xlarge": 0.899600, "r6id.16xlarge": 1.197000, "r6id.24xlarge": 1.795500, "r6id.2xlarge": 0.226400,
	"r6id.32xlarge": 2.394100, "r6id.4xlarge": 0.520100, "r6id.8xlarge": 0.620000, "r6id.large": 0.037400,
	"r6id.metal": 2.394100, "r6id.xlarge": 0.078100,
	// t1 family
	"t1.micro": 0.010500,
	// t2 family
	"t2.2xlarge": 0.111400, "t2.large": 0.031100, "t2.medium": 0.018200, "t2.micro": 0.003500,
	"t2.small": 0.009100, "t2.xlarge": 0.062400,
	// t3 family
	"t3.2xlarge": 0.153700, "t3.large": 0.027500, "t3.medium": 0.014000, "t3.micro": 0.003100,
	"t3.nano": 0.001600, "t3.small": 0.008000, "t3.xlarge": 0.067800,
	// t3a family
	"t3a.2xlarge": 0.133500, "t3a.large": 0.030100, "t3a.medium": 0.013400, "t3a.micro": 0.003900,
	"t3a.nano": 0.002300, "t3a.small": 0.007800, "t3a.xlarge": 0.061400,
	// t4g family
	"t4g.2xlarge": 0.080600, "t4g.large": 0.020200, "t4g.medium": 0.012600, "t4g.micro": 0.002500,
	"t4g.nano": 0.001400, "t4g.small": 0.005900, "t4g.xlarge": 0.040300,
	// vt1 family
	"vt1.24xlarge": 1.914400, "vt1.3xlarge": 0.195000, "vt1.6xlarge": 0.507300,
	// x1 family
	"x1.16xlarge": 2.000700, "x1.32xlarge": 4.001400,
	// x1e family
	"x1e.16xlarge": 4.003200, "x1e.2xlarge": 0.500400, "x1e.32xlarge": 8.006400, "x1e.4xlarge": 1.000800,
	"x1e.8xlarge": 2.001600, "x1e.xlarge": 0.250200,
	// x2gd family
	"x2gd.12xlarge": 4.008000, "x2gd.16xlarge": 5.344000, "x2gd.2xlarge": 0.200400, "x2gd.4xlarge": 0.408600,
	"x2gd.8xlarge": 0.801600, "x2gd.large": 0.050100, "x2gd.medium": 0.025100, "x2gd.metal": 5.344000,
	"x2gd.xlarge": 0.102500,
	// x2idn family
	"x2idn.16xlarge": 2.000700, "x2idn.24xlarge": 3.001100, "x2idn.32xlarge": 4.899000,
	"x2idn.metal": 6.303400,
	// x2iedn family
	"x2iedn.16xlarge": 4.001400, "x2iedn.24xlarge": 20.007000, "x2iedn.2xlarge": 0.500200,
	"x2iedn.32xlarge": 8.002800, "x2iedn.4xlarge": 1.000400, "x2iedn.8xlarge": 2.000700,
	"x2iedn.metal": 26.676000, "x2iedn.xlarge": 0.250100,
	// x2iezn family
	"x2iezn.12xlarge": 3.002400, "x2iezn.2xlarge": 0.500400, "x2iezn.4xlarge": 1.000800,
	"x2iezn.6xlarge": 1.501200, "x2iezn.8xlarge": 2.001600, "x2iezn.metal": 3.002400,
	// z1d family
	"z1d.12xlarge": 1.339200, "z1d.2xlarge": 0.286900, "z1d.3xlarge": 0.398600, "z1d.6xlarge": 0.669600,
	"z1d.large": 0.055800, "z1d.metal": 1.339200, "z1d.xlarge": 0.111600,
}
