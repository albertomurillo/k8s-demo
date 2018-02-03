FROM scratch
LABEL maintainer="albertomurillosilva@gmail.com"
ADD build/k8sdemo-backend /
CMD ["/k8sdemo-backend"]
