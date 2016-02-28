#include "mainwindow.h"
#include <QApplication>

int main(int argc, char *argv[])
{
    QApplication a(argc, argv);
    MainWindow w;
    w.SetQRCodeInfo(argv[1],argv[2]);
    w.show();

    return a.exec();
}
