FROM gcr.io/distroless/base-debian11
COPY axa-notebook-ssh /
ENTRYPOINT ["/axa-notebook-ssh"]
