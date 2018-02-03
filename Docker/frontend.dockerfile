FROM scratch
LABEL maintainer="albertomurillosilva@gmail.com"
ADD build/k8sdemo-frontend /
CMD ["/k8sdemo-frontend"]
