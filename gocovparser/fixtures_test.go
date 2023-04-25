package gocovparser_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

const repoName = "github.com/heynemann/go-cov-parser"

// CoverageFixture is a mock coverage.out contents.
func CoverageFixture(t *testing.T) string {
	t.Helper()

	return `
mode: set
github.com/heynemann/go-cov-parser/gocovparser/core.go:38.53,42.2 2 1
github.com/heynemann/go-cov-parser/gocovparser/core.go:45.60,47.20 2 1
github.com/heynemann/go-cov-parser/gocovparser/core.go:51.2,53.30 2 1
github.com/heynemann/go-cov-parser/gocovparser/core.go:110.2,110.22 1 1
github.com/heynemann/go-cov-parser/gocovparser/core.go:47.20,49.3 1 1
github.com/heynemann/go-cov-parser/gocovparser/core.go:53.30,55.22 2 1
github.com/heynemann/go-cov-parser/gocovparser/core.go:59.3,65.17 6 1
github.com/heynemann/go-cov-parser/gocovparser/core.go:69.3,70.17 2 1
github.com/heynemann/go-cov-parser/gocovparser/core.go:74.3,75.17 2 1
github.com/heynemann/go-cov-parser/gocovparser/core.go:79.3,80.17 2 1
github.com/heynemann/go-cov-parser/gocovparser/core.go:84.3,85.17 2 1
github.com/heynemann/go-cov-parser/gocovparser/core.go:89.3,107.4 2 1
github.com/heynemann/go-cov-parser/gocovparser/core.go:55.22,57.4 1 1
github.com/heynemann/go-cov-parser/gocovparser/core.go:65.17,67.4 1 0
github.com/heynemann/go-cov-parser/gocovparser/core.go:70.17,72.4 1 0
github.com/heynemann/go-cov-parser/gocovparser/core.go:75.17,77.4 1 0
github.com/heynemann/go-cov-parser/gocovparser/core.go:80.17,82.4 1 0
github.com/heynemann/go-cov-parser/gocovparser/core.go:85.17,87.4 1 0
github.com/heynemann/go-cov-parser/gocovparser/core.go:113.43,116.29 2 1
github.com/heynemann/go-cov-parser/gocovparser/core.go:130.2,130.15 1 1
github.com/heynemann/go-cov-parser/gocovparser/core.go:116.29,119.17 2 1
github.com/heynemann/go-cov-parser/gocovparser/core.go:123.3,123.39 1 1
github.com/heynemann/go-cov-parser/gocovparser/core.go:127.3,127.32 1 1
github.com/heynemann/go-cov-parser/gocovparser/core.go:119.17,120.12 1 1
github.com/heynemann/go-cov-parser/gocovparser/core.go:123.39,124.12 1 1
github.com/heynemann/go-cov-parser/gocovparser/core.go:133.40,135.16 2 1
github.com/heynemann/go-cov-parser/gocovparser/core.go:139.2,139.25 1 1
github.com/heynemann/go-cov-parser/gocovparser/core.go:135.16,137.3 1 0
github.com/heynemann/go-cov-parser/gocovparser/core.go:143.85,149.31 4 1
github.com/heynemann/go-cov-parser/gocovparser/core.go:186.2,186.20 1 1
github.com/heynemann/go-cov-parser/gocovparser/core.go:149.31,150.45 1 1
github.com/heynemann/go-cov-parser/gocovparser/core.go:154.3,154.49 1 1
github.com/heynemann/go-cov-parser/gocovparser/core.go:158.3,158.46 1 1
github.com/heynemann/go-cov-parser/gocovparser/core.go:162.3,162.29 1 1
github.com/heynemann/go-cov-parser/gocovparser/core.go:172.3,172.43 1 1
github.com/heynemann/go-cov-parser/gocovparser/core.go:150.45,152.4 1 1
github.com/heynemann/go-cov-parser/gocovparser/core.go:154.49,156.4 1 1
github.com/heynemann/go-cov-parser/gocovparser/core.go:158.46,160.4 1 1
github.com/heynemann/go-cov-parser/gocovparser/core.go:162.29,167.21 3 1
github.com/heynemann/go-cov-parser/gocovparser/core.go:167.21,169.5 1 1
github.com/heynemann/go-cov-parser/gocovparser/core.go:172.43,176.18 3 1
github.com/heynemann/go-cov-parser/gocovparser/core.go:182.4,182.59 1 1
github.com/heynemann/go-cov-parser/gocovparser/core.go:176.18,179.13 2 0
`
}

func CoverageFixture2(t *testing.T) string {
	t.Helper()

	return `
mode: atomic
github.cbhq.net/engineering/mongofle/crypt.go:29.115,36.2 1 28
github.cbhq.net/engineering/mongofle/crypt.go:39.107,43.50 1 69
github.cbhq.net/engineering/mongofle/crypt.go:47.2,48.55 2 67
github.cbhq.net/engineering/mongofle/crypt.go:52.2,53.16 2 67
github.cbhq.net/engineering/mongofle/crypt.go:57.2,58.16 2 67
github.cbhq.net/engineering/mongofle/crypt.go:62.2,67.16 1 67
github.cbhq.net/engineering/mongofle/crypt.go:71.2,72.16 2 64
github.cbhq.net/engineering/mongofle/crypt.go:76.2,76.20 1 64
github.cbhq.net/engineering/mongofle/crypt.go:43.50,45.3 1 2
github.cbhq.net/engineering/mongofle/crypt.go:48.55,50.3 1 0
github.cbhq.net/engineering/mongofle/crypt.go:53.16,55.3 1 0
github.cbhq.net/engineering/mongofle/crypt.go:58.16,60.3 1 0
github.cbhq.net/engineering/mongofle/crypt.go:67.16,69.3 1 3
github.cbhq.net/engineering/mongofle/crypt.go:72.16,74.3 1 0
github.cbhq.net/engineering/mongofle/crypt.go:80.104,84.50 1 64
github.cbhq.net/engineering/mongofle/crypt.go:88.2,89.63 2 64
github.cbhq.net/engineering/mongofle/crypt.go:93.2,94.16 2 64
github.cbhq.net/engineering/mongofle/crypt.go:98.2,99.16 2 64
github.cbhq.net/engineering/mongofle/crypt.go:103.2,108.16 1 64
github.cbhq.net/engineering/mongofle/crypt.go:112.2,113.16 2 64
github.cbhq.net/engineering/mongofle/crypt.go:117.2,117.20 1 64
github.cbhq.net/engineering/mongofle/crypt.go:84.50,86.3 1 0
github.cbhq.net/engineering/mongofle/crypt.go:89.63,91.3 1 0
github.cbhq.net/engineering/mongofle/crypt.go:94.16,96.3 1 0
github.cbhq.net/engineering/mongofle/crypt.go:99.16,101.3 1 0
github.cbhq.net/engineering/mongofle/crypt.go:108.16,110.3 1 0
github.cbhq.net/engineering/mongofle/crypt.go:113.16,115.3 1 0
github.cbhq.net/engineering/mongofle/crypt.go:123.129,134.2 1 0
github.cbhq.net/engineering/mongofle/crypt.go:141.25,143.2 1 0
github.cbhq.net/engineering/mongofle/crypt.go:146.105,148.2 1 0
github.cbhq.net/engineering/mongofle/crypt.go:151.26,152.2 0 28
github.cbhq.net/engineering/mongofle/crypt.go:155.45,157.2 1 161
github.cbhq.net/engineering/mongofle/default_mongo_encrypter.go:33.26,40.2 1 131
github.cbhq.net/engineering/mongofle/default_mongo_encrypter.go:43.108,52.16 4 55
github.cbhq.net/engineering/mongofle/default_mongo_encrypter.go:56.2,57.16 2 55
github.cbhq.net/engineering/mongofle/default_mongo_encrypter.go:61.2,62.16 2 55
github.cbhq.net/engineering/mongofle/default_mongo_encrypter.go:66.2,72.8 3 55
github.cbhq.net/engineering/mongofle/default_mongo_encrypter.go:52.16,54.3 1 0
github.cbhq.net/engineering/mongofle/default_mongo_encrypter.go:57.16,59.3 1 0
github.cbhq.net/engineering/mongofle/default_mongo_encrypter.go:62.16,64.3 1 0
github.cbhq.net/engineering/mongofle/default_mongo_encrypter.go:81.16,93.16 5 32
github.cbhq.net/engineering/mongofle/default_mongo_encrypter.go:97.2,106.16 8 32
github.cbhq.net/engineering/mongofle/default_mongo_encrypter.go:110.2,111.16 2 32
github.cbhq.net/engineering/mongofle/default_mongo_encrypter.go:115.2,116.16 2 32
github.cbhq.net/engineering/mongofle/default_mongo_encrypter.go:120.2,121.16 2 32
github.cbhq.net/engineering/mongofle/default_mongo_encrypter.go:125.2,127.16 3 32
github.cbhq.net/engineering/mongofle/default_mongo_encrypter.go:131.2,131.27 1 32
github.cbhq.net/engineering/mongofle/default_mongo_encrypter.go:93.16,95.3 1 0
github.cbhq.net/engineering/mongofle/default_mongo_encrypter.go:106.16,108.3 1 0
github.cbhq.net/engineering/mongofle/default_mongo_encrypter.go:111.16,113.3 1 0
github.cbhq.net/engineering/mongofle/default_mongo_encrypter.go:116.16,118.3 1 0
github.cbhq.net/engineering/mongofle/default_mongo_encrypter.go:121.16,123.3 1 0
github.cbhq.net/engineering/mongofle/default_mongo_encrypter.go:127.16,129.3 1 0
github.cbhq.net/engineering/mongofle/key_provider.go:40.65,45.2 1 0
github.cbhq.net/engineering/mongofle/key_provider.go:47.84,57.2 5 0
github.cbhq.net/engineering/mongofle/key_provider.go:59.87,61.16 2 0
github.cbhq.net/engineering/mongofle/key_provider.go:64.2,64.16 1 0
github.cbhq.net/engineering/mongofle/key_provider.go:67.2,67.29 1 0
github.cbhq.net/engineering/mongofle/key_provider.go:61.16,63.3 1 0
github.cbhq.net/engineering/mongofle/key_provider.go:64.16,66.3 1 0
github.cbhq.net/engineering/mongofle/key_provider.go:70.67,75.12 4 0
github.cbhq.net/engineering/mongofle/key_provider.go:78.2,78.33 1 0
github.cbhq.net/engineering/mongofle/key_provider.go:75.12,77.3 1 0
github.cbhq.net/engineering/mongofle/key_provider.go:81.107,85.34 3 0
github.cbhq.net/engineering/mongofle/key_provider.go:90.2,90.83 1 0
github.cbhq.net/engineering/mongofle/key_provider.go:85.34,86.24 1 0
github.cbhq.net/engineering/mongofle/key_provider.go:86.24,88.4 1 0
github.cbhq.net/engineering/mongofle/key_provider.go:114.24,119.16 2 28
github.cbhq.net/engineering/mongofle/key_provider.go:122.2,126.16 2 28
github.cbhq.net/engineering/mongofle/key_provider.go:130.2,136.8 1 28
github.cbhq.net/engineering/mongofle/key_provider.go:119.16,121.3 1 0
github.cbhq.net/engineering/mongofle/key_provider.go:126.16,128.3 1 0
github.cbhq.net/engineering/mongofle/key_provider.go:139.83,141.16 2 0
github.cbhq.net/engineering/mongofle/key_provider.go:145.2,147.64 3 0
github.cbhq.net/engineering/mongofle/key_provider.go:151.2,166.16 4 0
github.cbhq.net/engineering/mongofle/key_provider.go:170.2,175.17 3 0
github.cbhq.net/engineering/mongofle/key_provider.go:141.16,143.3 1 0
github.cbhq.net/engineering/mongofle/key_provider.go:147.64,149.3 1 0
github.cbhq.net/engineering/mongofle/key_provider.go:166.16,168.3 1 0
github.cbhq.net/engineering/mongofle/key_provider.go:178.86,180.15 2 131
github.cbhq.net/engineering/mongofle/key_provider.go:184.2,185.42 2 27
github.cbhq.net/engineering/mongofle/key_provider.go:189.2,190.16 2 27
github.cbhq.net/engineering/mongofle/key_provider.go:193.2,194.17 2 27
github.cbhq.net/engineering/mongofle/key_provider.go:180.15,182.3 1 104
github.cbhq.net/engineering/mongofle/key_provider.go:185.42,187.3 1 0
github.cbhq.net/engineering/mongofle/key_provider.go:190.16,192.3 1 0
github.cbhq.net/engineering/mongofle/key_provider.go:197.106,199.15 2 32
github.cbhq.net/engineering/mongofle/key_provider.go:203.2,208.42 3 0
github.cbhq.net/engineering/mongofle/key_provider.go:212.2,213.16 2 0
github.cbhq.net/engineering/mongofle/key_provider.go:216.2,217.17 2 0
github.cbhq.net/engineering/mongofle/key_provider.go:199.15,201.3 1 32
github.cbhq.net/engineering/mongofle/key_provider.go:208.42,210.3 1 0
github.cbhq.net/engineering/mongofle/key_provider.go:213.16,215.3 1 0
github.cbhq.net/engineering/mongofle/key_provider.go:220.80,221.25 1 27
github.cbhq.net/engineering/mongofle/key_provider.go:225.2,227.16 3 27
github.cbhq.net/engineering/mongofle/key_provider.go:231.2,234.16 4 27
github.cbhq.net/engineering/mongofle/key_provider.go:238.2,243.8 3 27
github.cbhq.net/engineering/mongofle/key_provider.go:221.25,223.3 1 0
github.cbhq.net/engineering/mongofle/key_provider.go:227.16,229.3 1 0
github.cbhq.net/engineering/mongofle/key_provider.go:234.16,236.3 1 0
github.cbhq.net/engineering/mongofle/key_provider.go:246.63,250.2 3 131
github.cbhq.net/engineering/mongofle/key_provider.go:252.82,255.30 3 32
github.cbhq.net/engineering/mongofle/key_provider.go:258.2,258.35 1 32
github.cbhq.net/engineering/mongofle/key_provider.go:255.30,257.3 1 0
github.cbhq.net/engineering/mongofle/key_provider.go:261.79,264.30 3 27
github.cbhq.net/engineering/mongofle/key_provider.go:267.2,268.12 2 27
github.cbhq.net/engineering/mongofle/key_provider.go:264.30,266.3 1 27
github.cbhq.net/engineering/mongofle/key_provider.go:268.12,270.3 1 27
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:19.119,21.50 1 72
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:26.2,26.31 1 71
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:30.2,31.40 2 70
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:35.2,38.33 2 70
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:43.2,43.33 1 53
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:48.2,48.12 1 40
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:21.50,23.3 1 1
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:26.31,28.3 1 1
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:31.40,33.3 1 154
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:38.33,40.3 1 17
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:43.33,45.3 1 13
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:56.9,58.16 2 17
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:62.2,62.27 1 17
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:73.2,73.12 1 16
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:58.16,60.3 1 0
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:62.27,64.17 2 17
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:68.3,68.94 1 16
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:64.17,66.4 1 1
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:68.94,70.4 1 0
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:76.92,77.30 1 17
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:85.2,85.39 1 1
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:77.30,80.41 2 19
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:80.41,82.4 1 16
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:88.101,89.27 1 13
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:99.2,99.32 1 3
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:120.2,120.39 1 3
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:89.27,92.41 2 15
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:92.41,93.71 1 10
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:93.71,95.5 1 10
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:99.32,100.38 1 3
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:100.38,102.18 2 3
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:106.4,106.35 1 3
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:102.18,104.5 1 0
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:106.35,109.43 2 6
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:113.5,113.72 1 0
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:109.43,110.14 1 6
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:113.72,115.6 1 0
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:123.52,126.40 2 10
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:137.2,137.15 1 10
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:127.14,128.25 1 0
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:129.14,130.26 1 0
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:131.14,132.28 1 7
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:133.10,134.19 1 3
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:146.9,148.30 1 43
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:196.2,196.12 1 43
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:148.30,152.46 3 118
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:164.3,164.63 1 53
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:179.3,179.63 1 42
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:152.46,154.18 2 65
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:158.4,160.12 2 65
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:154.18,156.5 1 0
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:164.63,171.18 1 11
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:175.4,175.12 1 11
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:171.18,173.5 1 0
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:179.63,180.32 1 2
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:180.32,181.67 1 6
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:181.67,188.20 1 6
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:188.20,190.7 1 0
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:204.9,206.16 2 13
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:210.2,210.27 1 13
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:236.2,236.12 1 10
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:206.16,208.3 1 0
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:210.27,212.17 2 13
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:216.3,217.17 2 13
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:221.3,222.17 2 13
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:226.3,227.17 2 10
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:231.3,231.94 1 10
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:212.17,214.4 1 0
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:217.17,219.4 1 0
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:222.17,224.4 1 3
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:227.17,229.4 1 0
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:231.94,233.4 1 0
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:239.66,243.16 3 30
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:247.2,247.18 1 30
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:251.2,251.27 1 30
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:262.2,262.20 1 30
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:243.16,245.3 1 0
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:247.18,249.3 1 0
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:251.27,255.17 3 30
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:259.3,259.31 1 30
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:255.17,257.4 1 0
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:265.59,274.12 4 97
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:278.2,279.16 2 97
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:283.2,283.20 1 96
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:274.12,276.3 1 0
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:279.16,281.3 1 1
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:286.41,288.16 2 18
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:292.2,292.17 1 18
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:288.16,290.3 1 0
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:296.114,298.21 1 69
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:303.2,303.31 1 69
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:307.2,307.49 1 68
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:311.2,311.12 1 52
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:298.21,300.3 1 0
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:303.31,305.3 1 1
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:307.49,309.3 1 16
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:314.90,315.35 1 15
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:323.2,323.17 1 1
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:315.35,318.41 2 19
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:318.41,320.4 1 14
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:331.9,333.16 2 16
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:337.2,338.16 2 15
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:342.2,342.35 1 15
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:362.2,362.12 1 14
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:333.16,335.3 1 1
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:338.16,340.3 1 0
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:342.35,344.17 2 15
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:348.3,349.17 2 15
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:353.3,353.28 1 15
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:357.3,357.87 1 14
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:344.17,346.4 1 0
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:349.17,351.4 1 0
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:353.28,355.4 1 1
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:357.87,359.4 1 0
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:371.9,373.30 1 32
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:418.2,418.12 1 32
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:373.30,376.68 2 90
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:386.3,386.63 1 90
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:401.3,401.63 1 78
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:376.68,378.18 2 38
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:382.4,382.29 1 38
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:378.18,380.5 1 0
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:386.63,393.18 1 12
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:397.4,397.12 1 12
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:393.18,395.5 1 0
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:401.63,402.32 1 2
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:402.32,403.67 1 6
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:403.67,410.20 1 6
github.cbhq.net/engineering/mongofle/mongo_encrypter.go:410.20,412.7 1 0
`
}

func CoverageFixture3(t *testing.T) string {
	t.Helper()

	contents, err := os.ReadFile("./coverage-fixture1.out")
	require.NoError(t, err)

	return string(contents)
}

func CoverageFixture4(t *testing.T) string {
	t.Helper()

	contents, err := os.ReadFile("./coverage-fixture2.out")
	require.NoError(t, err)

	return string(contents)
}

// CoverageFixture4 is a mock coverage.out contents.
func CoverageFixture5(t *testing.T) string {
	t.Helper()

	return `
mode: set
go.uber.org/zap/writer.go:50.65,52.16 2 1
`
}

func EmptyFixture(t *testing.T) string {
	t.Helper()

	return `
mode: set
`
}
