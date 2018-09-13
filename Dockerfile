FROM taion809/kafka-cli:1.0.1

VOLUME ["/configuration"]

COPY ensure /ensure

CMD /ensure