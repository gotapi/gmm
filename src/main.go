package main

import "C"
import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"net"
	"net/http"
	"strings"
)

/*
 #cgo LDFLAGS: -L/usr/local/lib/ -lmaxminddb
//
#include <maxminddb.h>
#include<stdio.h>
#include<string.h>
#include<stdlib.h>
#define ENTRY_TYPE_UINT32	0
#define ENTRY_TYPE_STRING	1
#define ENTRY_TYPE_DOUBLE	2
#define ENTRY_TYPE_BOOLEAN	3

#define COUNTRIES 250

char *country_code_array[] = {
	"AD", "AE", "AF", "AG", "AI", "AL", "AM", "AO", "AQ", "AR", "AS", "AT",
	"AU", "AW", "AX", "AZ", "BA", "BB", "BD", "BE", "BF", "BG", "BH", "BI",
	"BJ", "BL", "BM", "BN", "BO", "BQ", "BR", "BS", "BT", "BV", "BW", "BY",
	"BZ", "CA", "CC", "CD", "CF", "CG", "CH", "CI", "CK", "CL", "CM", "CN",
	"CO", "CR", "CU", "CV", "CW", "CX", "CY", "CZ", "DE", "DJ", "DK", "DM",
	"DO", "DZ", "EC", "EE", "EG", "EH", "ER", "ES", "ET", "FI", "FJ", "FK",
	"FM", "FO", "FR", "GA", "GB", "GD", "GE", "GF", "GG", "GH", "GI", "GL",
	"GM", "GN", "GP", "GQ", "GR", "GS", "GT", "GU", "GW", "GY", "HK", "HM",
	"HN", "HR", "HT", "HU", "ID", "IE", "IL", "IM", "IN", "IO", "IQ", "IR",
	"IS", "IT", "JE", "JM", "JO", "JP", "KE", "KG", "KH", "KI", "KM", "KN",
	"KP", "KR", "KW", "KY", "KZ", "LA", "LB", "LC", "LI", "LK", "LR", "LS",
	"LT", "LU", "LV", "LY", "MA", "MC", "MD", "ME", "MF", "MG", "MH", "MK",
	"ML", "MM", "MN", "MO", "MP", "MQ", "MR", "MS", "MT", "MU", "MV", "MW",
	"MX", "MY", "MZ", "NA", "NC", "NE", "NF", "NG", "NI", "NL", "NO", "NP",
	"NR", "NU", "NZ", "OM", "PA", "PE", "PF", "PG", "PH", "PK", "PL", "PM",
	"PN", "PR", "PS", "PT", "PW", "PY", "QA", "RE", "RO", "RS", "RU", "RW",
	"SA", "SB", "SC", "SD", "SE", "SG", "SH", "SI", "SJ", "SK", "SL", "SM",
	"SN", "SO", "SR", "SS", "ST", "SV", "SX", "SY", "SZ", "TC", "TD", "TF",
	"TG", "TH", "TJ", "TK", "TL", "TM", "TN", "TO", "TR", "TT", "TV", "TW",
	"TZ", "UA", "UG", "UM", "US", "UY", "UZ", "VA", "VC", "VE", "VG", "VI",
	"VN", "VU", "WF", "WS", "XK", "YE", "YT", "ZA", "ZM", "ZW"
};

char *country_code3_array[] = {
	"AND", "ARE", "AFG", "ATG", "AIA", "ALB", "ARM", "AGO", "ATA", "ARG",
	"ASM", "AUT", "AUS", "ABW", "ALA", "AZE", "BIH", "BRB", "BGD", "BEL",
	"BFA", "BGR", "BHR", "BDI", "BEN", "BLM", "BMU", "BRN", "BOL", "BES",
	"BRA", "BHS", "BTN", "BVT", "BWA", "BLR", "BLZ", "CAN", "CCK", "COD",
	"CAF", "COG", "CHE", "CIV", "COK", "CHL", "CMR", "CHN", "COL", "CRI",
	"CUB", "CPV", "CUW", "CXR", "CYP", "CZE", "DEU", "DJI", "DNK", "DMA",
	"DOM", "DZA", "ECU", "EST", "EGY", "ESH", "ERI", "ESP", "ETH", "FIN",
	"FJI", "FLK", "FSM", "FRO", "FRA", "GAB", "GBR", "GRD", "GEO", "GUF",
	"GGY", "GHA", "GIB", "GRL", "GMB", "GIN", "GLP", "GNQ", "GRC", "SGS",
	"GTM", "GUM", "GNB", "GUY", "HKG", "HMD", "HND", "HRV", "HTI", "HUN",
	"IDN", "IRL", "ISR", "IMN", "IND", "IOT", "IRQ", "IRN", "ISL", "ITA",
	"JEY", "JAM", "JOR", "JPN", "KEN", "KGZ", "KHM", "KIR", "COM", "KNA",
	"PRK", "KOR", "KWT", "CYM", "KAZ", "LAO", "LBN", "LCA", "LIE", "LKA",
	"LBR", "LSO", "LTU", "LUX", "LVA", "LBY", "MAR", "MCO", "MDA", "MNE",
	"MAF", "MDG", "MHL", "MKD", "MLI", "MMR", "MNG", "MAC", "MNP", "MTQ",
	"MRT", "MSR", "MLT", "MUS", "MDV", "MWI", "MEX", "MYS", "MOZ", "NAM",
	"NCL", "NER", "NFK", "NGA", "NIC", "NLD", "NOR", "NPL", "NRU", "NIU",
	"NZL", "OMN", "PAN", "PER", "PYF", "PNG", "PHL", "PAK", "POL", "SPM",
	"PCN", "PRI", "PSE", "PRT", "PLW", "PRY", "QAT", "REU", "ROU", "SRB",
	"RUS", "RWA", "SAU", "SLB", "SYC", "SDN", "SWE", "SGP", "SHN", "SVN",
	"SJM", "SVK", "SLE", "SMR", "SEN", "SOM", "SUR", "SSD", "STP", "SLV",
	"SXM", "SYR", "SWZ", "TCA", "TCD", "ATF", "TGO", "THA", "TJK", "TKL",
	"TLS", "TKM", "TUN", "TON", "TUR", "TTO", "TUV", "TWN", "TZA", "UKR",
	"UGA", "UMI", "USA", "URY", "UZB", "VAT", "VCT", "VEN", "VGB", "VIR",
	"VNM", "VUT", "WLF", "WSM", "XKX", "YEM", "MYT", "ZAF", "ZMB", "ZWE"
};

//
MMDB_s * telize_city;
MMDB_s * telize_asn;
int openMM(char * cityFile,char * asnFile){
		telize_city = malloc(sizeof(MMDB_s));
		telize_asn = malloc(sizeof(MMDB_s));
    if (MMDB_open(cityFile,MMDB_MODE_MMAP, telize_city) != MMDB_SUCCESS){
		printf("open cityFile failed:%s\n",cityFile);
		return 0;
	}

    if (MMDB_open(asnFile,MMDB_MODE_MMAP, telize_asn) != MMDB_SUCCESS) {
		printf("open asnFile failed:%s\n",asnFile);
		return 0;
	}
	printf("open asn file succ!\n");
    return 1;
}
void buf_append_int(char * json,char * field,uint32_t intValue){
    char dest[128]={""};
    sprintf(dest,",\"%s\":%d",field,intValue);
    strcat(json,dest);
}
void buf_append_double(char * json,char * field,double doubleValue){
    char dest[128]={""};
    sprintf(dest,",\"%s\":%.4f",field,doubleValue);
    strcat(json,dest);
}
void buf_append_boolean(char * json,char * field,bool boolValue){
    char dest[128]={""};
    sprintf(dest,",\"%s\":%s",field,boolValue?"true":"false");
    strcat(json,dest);
}
void buf_append_string(char *json,char *field,uint32_t data_size,const char * utf8string){
    char* strTemp = strndup(utf8string,data_size);
    char dest[256]={""};
    sprintf(dest,",\"%s\":\"%s\"",field,strTemp);
    strcat(json,dest);
}
void buf_append_chars(char *json,char *field,const char * chars){
    char dest[256]={""};
    sprintf(dest,",\"%s\":\"%s\"",field,chars);
    strcat(json,dest);
}
void telize_getdata(char *json, MMDB_lookup_result_s *lookup,
	MMDB_entry_data_s *entry_data, char *field, int type, ...)
{
    va_list keys;
    va_start(keys, type);

    MMDB_vget_value(&lookup->entry, entry_data, keys);

    if (entry_data->has_data) {
	switch(type) {
	    case ENTRY_TYPE_UINT32:
		buf_append_int(json,field, entry_data->uint32);
		break;
	    case ENTRY_TYPE_STRING:
		printf("data size:%d\n",entry_data->data_size);
		buf_append_string(json,field, entry_data->data_size,
			entry_data->utf8_string);
		break;
	    case ENTRY_TYPE_DOUBLE:
		buf_append_double(json,field, entry_data->double_value);
		break;
	    case ENTRY_TYPE_BOOLEAN:
		buf_append_boolean(json,field, entry_data->boolean);
		break;
	}
    }

    va_end(keys);
}
char * query( char * ip){
    int			slen, gai_error, mmdb_error;
    MMDB_lookup_result_s	lookup;
    MMDB_entry_data_s	entry_data;
    size_t tz_len;
	char *json = malloc(1024*sizeof(char));
	memset(json,0,1024);
 	strcat(json,"{\"ip\":\"");
    strcat(json,ip);
    strcat(json,"\"");

	    char *tz;
    lookup = MMDB_lookup_string(telize_city, ip, &gai_error, &mmdb_error);

    telize_getdata(json, &lookup, &entry_data,
	    "continent_code", ENTRY_TYPE_STRING, "continent", "code", NULL);
    telize_getdata(json, &lookup, &entry_data,
	    "country", ENTRY_TYPE_STRING, "country", "names", "en", NULL);
    telize_getdata(json, &lookup, &entry_data, "is_in_european_union",
	    ENTRY_TYPE_BOOLEAN, "country", "is_in_european_union", NULL);

    telize_getdata(json, &lookup, &entry_data, "region",
	    ENTRY_TYPE_STRING, "subdivisions", "0", "names", "en", NULL);

    telize_getdata(json, &lookup, &entry_data, "region_code",
	    ENTRY_TYPE_STRING, "subdivisions", "0", "iso_code", NULL);

    telize_getdata(json, &lookup, &entry_data, "city",
	    ENTRY_TYPE_STRING, "city", "names", "en", NULL);

    telize_getdata(json, &lookup, &entry_data, "postal_code",
	    ENTRY_TYPE_STRING, "postal", "code", NULL);

    telize_getdata(json, &lookup, &entry_data, "latitude",
	    ENTRY_TYPE_DOUBLE, "location", "latitude", NULL);

    telize_getdata(json, &lookup, &entry_data, "longitude",
	    ENTRY_TYPE_DOUBLE, "location", "longitude", NULL);

    MMDB_get_value(&lookup.entry, &entry_data, "country", "iso_code", NULL);
	if (entry_data.has_data) {
		buf_append_string(json, "country_code",
		    entry_data.data_size, entry_data.utf8_string);
		for (size_t loop = 0; loop < COUNTRIES; loop++) {
			if (!strncmp(country_code_array[loop],
			    entry_data.utf8_string, 2)) {
				buf_append_chars(json,"country_code3",country_code3_array[loop]);
				break;
			}
		}
	}
	MMDB_get_value(&lookup.entry, &entry_data,"location", "time_zone", NULL);
	if (entry_data.has_data) {
        		tz_len = entry_data.data_size;
        		tz = strndup(entry_data.utf8_string, tz_len);
        		if (tz) {
        			buf_append_chars(json,"tz",tz);
        		}
    }

    lookup = MMDB_lookup_string(telize_asn, ip, &gai_error, &mmdb_error);
    telize_getdata(json, &lookup, &entry_data, "asn",
	    ENTRY_TYPE_UINT32, "autonomous_system_number", NULL);

    telize_getdata(json, &lookup, &entry_data, "organization",
	    ENTRY_TYPE_STRING, "autonomous_system_organization", NULL);

    strcat(json,"}\n");
	return json;

}
void closeMM(){
 	MMDB_close(telize_city);
    MMDB_close(telize_asn);
}
*/
import "C"

var privateIPBlocks []*net.IPNet

type IpGeo struct {
	Ip            string  `json:"ip"`
	ContinentCode string  `json:"continent_code"`
	Country       string  `json:"country"`
	Latitude      float64 `json:"latitude"`
	Longitude     float64 `json:"longitude"`
	CountryCode   string  `json:"country_code"`
	CountryCode3  string  `json:"country_code3"`
	Tz            string  `json:"string"`
	Asn           int     `json:"asn"`
	Organization  string  `json:"organization"`
}

func Init() {
	for _, cidr := range []string{
		"127.0.0.0/8",    // IPv4 loopback
		"10.0.0.0/8",     // RFC1918
		"172.16.0.0/12",  // RFC1918
		"192.168.0.0/16", // RFC1918
		"169.254.0.0/16", // RFC3927 link-local
		"::1/128",        // IPv6 loopback
		"fe80::/10",      // IPv6 link-local
		"fc00::/7",       // IPv6 unique local addr
	} {
		_, block, err := net.ParseCIDR(cidr)
		if err != nil {
			panic(fmt.Errorf("parse error on %q: %v", cidr, err))
		}
		privateIPBlocks = append(privateIPBlocks, block)
	}
}

func query(ip string) (*IpGeo, error) {
	str := C.GoString(C.query(C.CString(ip)))
	var geo IpGeo
	er := json.Unmarshal([]byte(str), &geo)
	if er != nil {
		return &IpGeo{}, er
	}
	return &geo, nil
}

func isPrivateIP(ip net.IP) bool {
	if ip.IsLoopback() || ip.IsLinkLocalUnicast() || ip.IsLinkLocalMulticast() {
		return true
	}
	for _, block := range privateIPBlocks {
		if block.Contains(ip) {
			return true
		}
	}
	return false
}
func getIp(r *http.Request) string {
	obj := r.Header.Values("x-real-ip")
	xRealIpStr := ""
	if len(obj) > 0 {
		idx := strings.LastIndex(obj[0], ":")
		if idx > 2 {
			xRealIpStr = obj[0][:idx]
		} else {
			xRealIpStr = obj[0]
		}
	} else {
		idx := strings.LastIndex(r.RemoteAddr, ":")
		xRealIpStr = r.RemoteAddr[:idx]
	}
	return xRealIpStr
}

func main() {
	var dbDir string
	flag.StringVar(&dbDir, "db_dir", "/var/db/GeoIP/", "directory path of maxmind database")
	asnFileLocation := dbDir + "GeoLite2-ASN.mmdb"
	cityFileLocation := dbDir + "GeoLite2-City.mmdb"
	cityFile := C.CString(cityFileLocation)
	asnFile := C.CString(asnFileLocation)
	openRes := C.openMM(cityFile, asnFile)
	if openRes != 1 {
		panic("can't open city/asn files")
		return
	}
	Init()

	defer func() {
		C.closeMM()
	}()

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/location/:ip", func(c *gin.Context) {
		var ip = c.Param("ip")
		geo, er := query(ip)

		if er != nil {
			c.JSON(500, er)
			return
		}
		c.JSON(200, geo)
	})
	r.GET("/location/", func(c *gin.Context) {
		ipString := getIp(c.Request)
		fmt.Printf("got ip:%s\n", ipString)
		ip := net.ParseIP(ipString)
		if ip == nil {
			c.JSON(500, gin.H{
				"message": "ip invalid",
			})
			return
		}
		if isPrivateIP(ip) {
			c.JSON(200, &IpGeo{Organization: "-", Asn: 0, Country: "-",
				Latitude: 0, Longitude: 0, CountryCode: "-", CountryCode3: "", Tz: "", Ip: ipString})
			return
		}
		geo, er := query(ipString)
		if er != nil {
			c.JSON(500, er)
			return
		}
		_, er = json.Marshal(geo)
		c.JSON(200, geo)

	})
	r.Run(":7654")

}
