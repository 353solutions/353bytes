FROM golang:1.14-buster
WORKDIR /code
RUN echo "export PS1='$ '" >> ~/.bashrc
RUN echo 'set -o vi' >> ~/.bashrc
RUN apt-get update && apt-get install -y vim curl
RUN curl -fLo ~/.vim/autoload/plug.vim --create-dirs \
    https://raw.githubusercontent.com/junegunn/vim-plug/master/plug.vim
COPY vimrc /root/.vimrc
RUN vim +PlugInstall +qa
RUN vim +GoUpdateBinaries +qa
