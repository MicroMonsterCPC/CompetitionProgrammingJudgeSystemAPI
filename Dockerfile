FROM ubuntu
RUN apt-get update\
			&& apt-get install -y ruby python clang time binutils npm python-software-properties
# RUN add-apt-repository ppa:directhex/monoxide

RUN mkdir /judge
WORKDIR /judge
COPY . /judge

CMD ["ruby","Setup.rb"]
