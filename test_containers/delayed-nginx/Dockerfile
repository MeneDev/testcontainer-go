FROM nginx:1.15.2

ADD delay.sh /delay.sh
RUN chmod +x /delay.sh

ENTRYPOINT "/delay.sh"