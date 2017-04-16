#include <iostream>
#include <regex>

extern "C" unsigned char cxxinit(char *);
extern "C" unsigned char cxxse(long long,char *);

using namespace std;
static regex re;
unsigned char cxxinit(char *arg0){
	try{
		re=regex(arg0);
		return 0;
	}catch(...){
	}
	cerr<<"error\n";
	return 1;
}
unsigned char cxxse(long long len,char *filecont){
	try{
		return regex_search(filecont,filecont+len,re);
	}catch(...){
	}
	cerr<<"error\n";
	return 2;
}
