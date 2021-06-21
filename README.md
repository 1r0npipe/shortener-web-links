# shortener-web-links
The service provides a short link for long one based on Go

Assuming the project is split into two sub-folders: backend and frontend, the REST API calls are used between those two layares.

- Backend is written on Go (+ some external storage approach?)
- Frontend is written on ?? (node.js/TypeScript?)

Those two parts are also will be using docker/k8s service (like minikube?) to deploy into, docker composer can be used as well.

The communication will be done under TCP/IP layout:
1. At Frontend user puts the long web-link and click "Generate", the short link is being provided
2. The information about long and short link goes into storage
3. Every time anyone uses the short link the information (statistic) gets updated inside of application
4. That statistic will be shown to admin/anyone who needs

P.S. this is initial proposition, some design could be changed in between.

# Choosing web-framework (router) for project
Current date: 10 Jun 2021
|Name (router)-web-framework| Stars on GitHub| Last release|Documentation|Searches on StackOverflow|Features|
|:---:|:---:|:---:|:---:|:---:|:---:|
|Gin|	    48.7k|	19 days ago|	    Excellent|	        272|Middleware support/Json validation/render builtIn/fast performance(?)|
|Beego|	    26.5k|	Nov 2020|	        Super excellent|	69|	Native Http Go package/monitor CPU perf|
|Echo|	    20k|	8 May 2021|	        Super  Excellent|	27|	Automatic TLS/smart routes/zero dynamic memory allocation|
|fasthttp|	15.2k|	9 days ago	|       Good|               16|	Optimize for low memory usage/fast performance(?)|
|mux (gorilla)|	14.5k|	Aug 2020|	    Good|	            151|Impements http.handler interface/subroutes|
|httprouter|	12.8|	Sep 2019|	    Poor|	            28| Zero garbage/multiple matching for route|

- "?" mark means there is no exact information and proof for that 
* Summary: Based on the information from that table, I will be using the Gin web-framework with all the intergrated features within, also its popularity and documentation (not ideal but still good enough) is really good which will be helping me to use that approach at any another project with same web-framework (router). Moreover it shows good result for search at StackOverflow, which means people meet and solve the problem frequently, also it is one of the routers (frameworks overall) which has last release just recently (19 days ago for current date). And last poing: I have worked with that approach already, a little, but would be good to proceed to get it known better.
