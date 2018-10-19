FROM autotool:v7.2

MAINTAINER zhuo.li-2@yeepay.com

ADD ./mybinary /root/hook/
ADD ./autotitletool /root/hook/
ADD ./maupassant /root/website/themes/maupassant
ADD ./config.toml /root/website
ENTRYPOINT ["/root/hook/mybinary"]