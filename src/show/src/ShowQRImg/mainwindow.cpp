#include "mainwindow.h"
#include "ui_mainwindow.h"

MainWindow::MainWindow(QWidget *parent) :
    QMainWindow(parent),
    ui(new Ui::MainWindow)
{
    ui->setupUi(this);
}

void MainWindow::SetQRCodeInfo(char *addrstr, char *filename)
{
    this->_filename = QString(filename);
    this->_addrinfo = QString(addrstr);
    if(this->_addrinfo.isNull()||this->_addrinfo.isEmpty()){
        QMessageBox::information(this,"错误","地址参数不正确!");
        return;
    }
    if(this->_filename.isNull()||this->_filename.isEmpty()){
        QMessageBox::information(this,"错误","没有找到图片!");
        return;
    }
    //read image
    QPixmap img(this->_filename);
    //show image
    ui->labelImg->setPixmap(img);
//auto size layout
    ui->labelImg->resize(img.width(),img.height());
    ui->labelInfo->setText(this->_addrinfo);
    ui->centralWidget->resize(ui->labelImg->width(),ui->labelImg->height()+ui->labelInfo->height()+10);

}

MainWindow::~MainWindow()
{
    delete ui;
}
