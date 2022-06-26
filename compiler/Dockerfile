# use the cached image from Dockerfile.deps
# to allow (re)building without internet access
FROM faust.cs.fau.de:5000/compiler60-compiler-deps

EXPOSE 6061

COPY ./src/main /app/src/main

RUN \
  # compile and generate parser 
  gradle build && \ 
  # to make the /bin/echo hack in getJavaArgs work
  ln -s /opt/java/openjdk/bin/java /bin/java && \
  # create start.sh with arguments from getJavaArgs
  # limit memory between 20MB-100MB  
  echo exec java -Xms20m -Xmx100m `gradle -q getJavaArgs` > start.sh && \
  chmod +x start.sh

CMD ["/bin/sh", "./start.sh"]
