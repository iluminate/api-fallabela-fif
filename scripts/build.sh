echo "generate image..."
docker build --file build/Dockerfile . -t kevin24ec/api-fallabela-fif:1.0

echo "upload image..."
docker push kevin24ec/api-fallabela-fif:1.0