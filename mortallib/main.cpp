
//curl 예제 출처[https://www.joinc.co.kr/w/Site/Web/documents/UsedCurl]
/*
	 sample for O'ReillyNet article on libcURL:
	 {TITLE}
	 {URL}
AUTHOR: Ethan McCallum

Scenario: use http/GET to fetch a webpage

이 코드는 Ubuntu 리눅스 Kernel 2.6.15에서 
libcURL 버젼 7.15.1로 테스트 되었다.
2006년 8월 3일
 */

#include <iostream>
#include <cstring>
#include "libyuncoil.h"



extern "C" {
#include<curl/curl.h>
}

// - - - - - - - - - - - - - - - - - - - -

enum {
	ERROR_ARGS = 1 ,
	ERROR_CURL_INIT = 2
} ;

enum {
	OPTION_FALSE = 0 ,
	OPTION_TRUE = 1
} ;

enum {
	FLAG_DEFAULT = 0 
} ;

// - - - - - - - - - - - - - - - - - - - -

int main( const int argc , const char** argv ){

#if 0
	if( argc != 2 ){
		std::cerr << " Usage: ./" << argv[0] << " {url} [debug]" << std::endl ;
		return( ERROR_ARGS ) ;
	}
#endif

	const char* url = "www.hacker.org/coil/index.php?gotolevel=8&go=Go+To+Level" ;

	// lubcURL 초기화 
	curl_global_init( CURL_GLOBAL_ALL ) ;

	// context객체의 생성
	CURL* ctx = curl_easy_init() ;

	if( NULL == ctx ){
		std::cerr << "Unable to initialize cURL interface" << std::endl ;
		return( ERROR_CURL_INIT ) ;
	}

	// context 객체를 설정한다.	
	// 긁어올 url을 명시하고, url이 URL정보임을 알려준다.
	curl_easy_setopt( ctx , CURLOPT_URL,  url ) ;

	//로그인한 coil-session
	curl_easy_setopt( ctx , CURLOPT_COOKIEFILE,  "conf/phpsession" ) ; //session cookie

	// no progress bar:
	curl_easy_setopt( ctx , CURLOPT_NOPROGRESS , OPTION_TRUE ) ;

	/*
		 By default, headers are stripped from the output.
		 They can be:

		 - passed through a separate FILE* (CURLOPT_WRITEHEADER)

		 - included in the body's output (CURLOPT_HEADER -> nonzero value)
		 (here, the headers will be passed to whatever function
		 processes the body, along w/ the body)

		 - handled with separate callbacks (CURLOPT_HEADERFUNCTION)
		 (in this case, set CURLOPT_WRITEHEADER to a
		 matching struct for the function)

	 */

	// 헤더는 표준에러로 출력하도록 하다. 
	curl_easy_setopt( ctx , CURLOPT_WRITEHEADER , stderr ) ;


	// body 데이터는 표준출력 하도록 한다.
	curl_easy_setopt( ctx , CURLOPT_WRITEDATA , stdout ) ;

	// context 객체의 설정 종료 


	// 웹페이지를 긁어온다. 

	const CURLcode rc = curl_easy_perform( ctx ) ;

	if( CURLE_OK != rc ){

		std::cerr << "Error from cURL: " << curl_easy_strerror( rc ) << std::endl ;

	}else{

		// get some info about the xfer:
		double statDouble ;
		long statLong ;
		char* statString = NULL ;

		// HTTP 응답코드를 얻어온다. 
		if( CURLE_OK == curl_easy_getinfo( ctx , CURLINFO_HTTP_CODE , &statLong ) ){
			std::cout << "Response code:  " << statLong << std::endl ;
		}

		// Content-Type 를 얻어온다.
		if( CURLE_OK == curl_easy_getinfo( ctx , CURLINFO_CONTENT_TYPE , &statString ) ){
			std::cout << "Content type:   " << statString << std::endl ;
		}

		// 다운로드한 문서의 크기를 얻어온다.
		if( CURLE_OK == curl_easy_getinfo( ctx , CURLINFO_SIZE_DOWNLOAD , &statDouble ) ){
			std::cout << "Download size:  " << statDouble << "bytes" << std::endl ;
		}

		// 
		if( CURLE_OK == curl_easy_getinfo( ctx , CURLINFO_SPEED_DOWNLOAD , &statDouble ) ){
			std::cout << "Download speed: " << statDouble << "bytes/sec" << std::endl ;
		}

	}

	// cleanup
	curl_easy_cleanup( ctx ) ;
	curl_global_cleanup() ;

	return( 0 ) ;

} // main()
