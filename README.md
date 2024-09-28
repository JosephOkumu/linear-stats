# linear-stats

# Description
This Go program reads data from a file and calculates linear regression and pearson correlation and prints them on the screen.

## Documentation
This section provides details on how to use linear-stats program.

### Installation
To use the this program, you need to have Go installed on your system. You can download and install the latest version of Go from [here](https://go.dev/doc/install).

### Usage
1. Clone this repository to your local machine using the following command:
    ```bash
    git clone https://github.com/JosephOkumu/linear-stats
    ```
2. Navigate into the linear-stats directory:
    ```bash
    cd linear-stats
    ```
3. Put your data in in a folder and name it data.txt, placd it in the linear-stats parent folder and run the  program using:
    ```bash
    go run . data.txt
    ```
    The output should be in the form: 

    ```bash
    Linear Regression Line: y = <value>x + <value>
    Pearson Correlation Coefficient: <value>
    ```

### Testing
1.  Download the zip file for testing [here](https://assets.01-edu.org/stats-projects/stat-bin-dockerized.zip)
2.  Extract the zip file.
3.  Open the stats-bin folder
4.  Copy the bin folder and paste it in your linear-stats root folder.
5.  Now run the command "./bin/linear-stats" and again run "go run main.go data.txt" and compare if the results match.

### Features
- Calculates linear regression.
- Calculates pearson correlation.

### Contributions
Pull requests are welcome. You can contribute to this project by adding new features, optimizing the algorithm, or fixing bugs.

For major changes, please open an issue first to discuss what you would like to change.

### Author
[JosephOKumu](https://github.com/JosephOkumu)

### License
[LICENSE](./LICENSE)


### Credits
[Zone01Kisumu](https://www.zone01kisumu.ke/)
