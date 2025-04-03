

#quarcus app

mvn io.quarkus:quarkus-maven-plugin:2.0.0.Final:create \
    -DprojectGroupId=mc.study.urlshortener \
    -DprojectArtifactId=quarkus-url-shortener \
    -DclassName="mc.study.urlshotener.api.UrlResource" \
    -Dpath="/hello"

./mvnw compile quarkus:dev


./mvnw package -Pnative
docker build -f src/main/docker/Dockerfile.native -t quarkus/quarkus-url-shortener .
