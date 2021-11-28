FROM scratch
ADD main /main
EXPOSE 8080
CMD  ["/main"]