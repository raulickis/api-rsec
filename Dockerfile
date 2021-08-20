FROM alpine
COPY main /
EXPOSE 9990
CMD cd / ; /main
