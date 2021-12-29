FROM centurylink/ca-certs
MAINTAINER info@wimark.com

ADD ./bin/manufacturer-parser /usr/bin/manufacturer-parser

ENTRYPOINT [ "/usr/bin/manufacturer-parser"]
