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