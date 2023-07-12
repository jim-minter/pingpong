FROM scratch
COPY pingpong pingpong
ENTRYPOINT [ "/pingpong" ]