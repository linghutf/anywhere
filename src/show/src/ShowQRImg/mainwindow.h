#ifndef MAINWINDOW_H
#define MAINWINDOW_H

#include <QMainWindow>
#include <QMessageBox>

namespace Ui {
class MainWindow;
}

class MainWindow : public QMainWindow
{
    Q_OBJECT

public:
    explicit MainWindow(QWidget *parent = 0);
    void SetQRCodeInfo(char *addstr,char *filename);

    ~MainWindow();

private:
    Ui::MainWindow *ui;
    QString _addrinfo;
    QString _filename;

};

#endif // MAINWINDOW_H
