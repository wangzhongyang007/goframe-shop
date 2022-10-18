FROM loads/alpine:3.8

LABEL maintainer="425772719@qq.com"

###############################################################################
#                                INSTALLATION
###############################################################################

# 设置固定的项目路径
ENV WORKDIR /go/src/shop

# 添加应用可执行文件，并设置执行权限
ADD ./shop   $WORKDIR/shop
RUN chmod +x $WORKDIR/shop

# 添加I18N多语言文件、静态文件、配置文件、模板文件
ADD i18n     $WORKDIR/i18n
ADD public   $WORKDIR/public
ADD config   $WORKDIR/config
ADD template $WORKDIR/template

###############################################################################
#                                   START
###############################################################################
WORKDIR $WORKDIR
CMD ./shop
