package main

import (
	"encoding/binary"
	"fmt"
	"math"
	"math/rand"
	"os"
	"time"
)

// DNN结构体定义
type DNN struct {
	inputSize    int
	hiddenSize1  int
	hiddenSize2  int
	outputSize   int
	learningRate float64
	weights1     [][]float64
	weights2     [][]float64
	weights3     [][]float64
}

// 激活函数和其导数
func relu(x float64) float64 {
	if x > 0 {
		return x
	}
	return 0
}

func reluDerivative(x float64) float64 {
	if x > 0 {
		return 1
	}
	return 0
}

func softmax(x []float64) []float64 {
	expSum := 0.0
	for i := range x {
		x[i] = math.Exp(x[i])
		expSum += x[i]
	}
	for i := range x {
		x[i] /= expSum
	}
	return x
}

// 加载MNIST数据集
func loadMNISTImages(filename string) ([][]float64, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var magic, num, rows, cols int32
	binary.Read(file, binary.BigEndian, &magic)
	binary.Read(file, binary.BigEndian, &num)
	binary.Read(file, binary.BigEndian, &rows)
	binary.Read(file, binary.BigEndian, &cols)

	images := make([][]float64, num)
	for i := 0; i < int(num); i++ {
		images[i] = make([]float64, rows*cols)
		for j := 0; j < int(rows*cols); j++ {
			var pixel uint8
			binary.Read(file, binary.BigEndian, &pixel)
			images[i][j] = float64(pixel) / 255.0
		}
	}
	return images, nil
}

func loadMNISTLabels(filename string) ([]int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var magic, num int32
	binary.Read(file, binary.BigEndian, &magic)
	binary.Read(file, binary.BigEndian, &num)

	labels := make([]int, num)
	for i := 0; i < int(num); i++ {
		var label uint8
		binary.Read(file, binary.BigEndian, &label)
		labels[i] = int(label)
	}
	return labels, nil
}

// 初始化权重
func initializeWeights(inputSize, outputSize int) [][]float64 {
	weights := make([][]float64, inputSize)
	for i := range weights {
		weights[i] = make([]float64, outputSize)
		for j := range weights[i] {
			weights[i][j] = rand.Float64()*2 - 1
		}
	}
	return weights
}

// DNN结构体的方法
func (dnn *DNN) forward(input []float64) ([]float64, []float64, []float64) {
	hidden1 := make([]float64, len(dnn.weights1[0]))
	for i := range hidden1 {
		for j := range input {
			hidden1[i] += input[j] * dnn.weights1[j][i]
		}
		hidden1[i] = relu(hidden1[i])
	}

	hidden2 := make([]float64, len(dnn.weights2[0]))
	for i := range hidden2 {
		for j := range hidden1 {
			hidden2[i] += hidden1[j] * dnn.weights2[j][i]
		}
		hidden2[i] = relu(hidden2[i])
	}

	output := make([]float64, len(dnn.weights3[0]))
	for i := range output {
		for j := range hidden2 {
			output[i] += hidden2[j] * dnn.weights3[j][i]
		}
	}
	output = softmax(output)
	return hidden1, hidden2, output
}

func (dnn *DNN) train(images [][]float64, labels []int, epochs int) {
	for epoch := 0; epoch < epochs; epoch++ {
		totalLoss := 0.0
		for i, input := range images {
			label := labels[i]

			// 前向传播
			hidden1, hidden2, output := dnn.forward(input)

			// 计算损失和误差
			target := make([]float64, dnn.outputSize)
			target[label] = 1.0
			outputError := make([]float64, dnn.outputSize)
			for j := range output {
				outputError[j] = target[j] - output[j]
				totalLoss += 0.5 * (target[j] - output[j]) * (target[j] - output[j])
			}

			hidden2Error := make([]float64, dnn.hiddenSize2)
			for j := range hidden2 {
				for k := range outputError {
					hidden2Error[j] += outputError[k] * dnn.weights3[j][k]
				}
				hidden2Error[j] *= reluDerivative(hidden2[j])
			}

			hidden1Error := make([]float64, dnn.hiddenSize1)
			for j := range hidden1 {
				for k := range hidden2Error {
					hidden1Error[j] += hidden2Error[k] * dnn.weights2[j][k]
				}
				hidden1Error[j] *= reluDerivative(hidden1[j])
			}

			// 反向传播和权重更新
			for j := range dnn.weights3 {
				for k := range dnn.weights3[j] {
					dnn.weights3[j][k] += dnn.learningRate * outputError[k] * hidden2[j]
				}
			}

			for j := range dnn.weights2 {
				for k := range dnn.weights2[j] {
					dnn.weights2[j][k] += dnn.learningRate * hidden2Error[k] * hidden1[j]
				}
			}

			for j := range dnn.weights1 {
				for k := range dnn.weights1[j] {
					dnn.weights1[j][k] += dnn.learningRate * hidden1Error[k] * input[j]
				}
			}
		}
		fmt.Printf("Epoch %d/%d, Loss: %f\n", epoch+1, epochs, totalLoss/float64(len(images)))
	}
}

func (dnn *DNN) predict(input []float64) int {
	_, _, output := dnn.forward(input)
	maxIndex := 0
	for i := range output {
		if output[i] > output[maxIndex] {
			maxIndex = i
		}
	}
	return maxIndex
}

func (dnn *DNN) evaluate(images [][]float64, labels []int) float64 {
	correct := 0
	for i, input := range images {
		prediction := dnn.predict(input)
		if prediction == labels[i] {
			correct++
		}
	}
	return float64(correct) / float64(len(labels))
}

// NewDNN 创建和初始化DNN
func NewDNN(inputSize, hiddenSize1, hiddenSize2, outputSize int, learningRate float64) *DNN {
	return &DNN{
		inputSize:    inputSize,
		hiddenSize1:  hiddenSize1,
		hiddenSize2:  hiddenSize2,
		outputSize:   outputSize,
		learningRate: learningRate,
		weights1:     initializeWeights(inputSize, hiddenSize1),
		weights2:     initializeWeights(hiddenSize1, hiddenSize2),
		weights3:     initializeWeights(hiddenSize2, outputSize),
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	trainImages, err := loadMNISTImages("train-images.idx3-ubyte")
	if err != nil {
		fmt.Println("Failed to load training images:", err)
		return
	}

	trainLabels, err := loadMNISTLabels("train-labels.idx1-ubyte")
	if err != nil {
		fmt.Println("Failed to load training labels:", err)
		return
	}

	testImages, err := loadMNISTImages("t10k-images.idx3-ubyte")
	if err != nil {
		fmt.Println("Failed to load test images:", err)
		return
	}

	testLabels, err := loadMNISTLabels("t10k-labels.idx1-ubyte")
	if err != nil {
		fmt.Println("Failed to load test labels:", err)
		return
	}

	epochs := 10
	learningRate := 0.01

	dnn := NewDNN(28*28, 128, 64, 10, learningRate)
	dnn.train(trainImages, trainLabels, epochs)

	accuracy := dnn.evaluate(testImages, testLabels)
	fmt.Printf("Model accuracy on test set: %.2f%%\n", accuracy*100)
}
