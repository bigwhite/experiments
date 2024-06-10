package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

// Helper function to read CSV files
func readCSV(filePath string) ([][]float64, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	data := make([][]float64, len(records)-1)
	for i := 1; i < len(records); i++ {
		data[i-1] = make([]float64, len(records[i]))
		for j := range records[i] {
			data[i-1][j], err = strconv.ParseFloat(records[i][j], 64)
			if err != nil {
				return nil, err
			}
		}
	}
	return data, nil
}

// Helper function to standardize data
func standardize(data [][]float64) ([][]float64, []float64, []float64) {
	mean := make([]float64, len(data[0])-1)
	std := make([]float64, len(data[0])-1)
	for i := 0; i < len(data[0])-1; i++ {
		for j := 0; j < len(data); j++ {
			mean[i] += data[j][i]
		}
		mean[i] /= float64(len(data))
	}
	for i := 0; i < len(data[0])-1; i++ {
		for j := 0; j < len(data); j++ {
			std[i] += math.Pow(data[j][i]-mean[i], 2)
		}
		std[i] = math.Sqrt(std[i] / float64(len(data)))
	}
	standardizedData := make([][]float64, len(data))
	for i := 0; i < len(data); i++ {
		standardizedData[i] = make([]float64, len(data[i]))
		for j := 0; j < len(data[i])-1; j++ {
			standardizedData[i][j] = (data[i][j] - mean[j]) / std[j]
		}
		standardizedData[i][len(data[i])-1] = data[i][len(data[i])-1]
	}
	return standardizedData, mean, std
}

// Neural Network Layer structure
type Layer struct {
	weights []float64
	bias    float64
}

// Initialize a layer with the given number of inputs
func NewLayer(inputSize int) *Layer {
	weights := make([]float64, inputSize)
	for i := range weights {
		weights[i] = 0.01 // small random values, here we use a small constant for simplicity
	}
	return &Layer{
		weights: weights,
		bias:    0.0,
	}
}

// Forward propagation
func (layer *Layer) Forward(inputs []float64) float64 {
	output := layer.bias
	for i := range layer.weights {
		output += layer.weights[i] * inputs[i]
	}
	return output
}

// Backward propagation (gradient computation and update)
func (layer *Layer) Backward(inputs []float64, error float64, learningRate float64) {
	for i := range layer.weights {
		layer.weights[i] -= learningRate * error * inputs[i]
	}
	layer.bias -= learningRate * error
}

// Training the neural network
func trainModel(data [][]float64, learningRate float64, epochs int) *Layer {
	features := len(data[0]) - 1
	layer := NewLayer(features)

	for epoch := 0; epoch < epochs; epoch++ {
		totalError := 0.0
		for i := 0; i < len(data); i++ {
			inputs := data[i][:features]
			target := data[i][features]
			prediction := layer.Forward(inputs)
			error := prediction - target
			totalError += error * error
			layer.Backward(inputs, error, learningRate)
		}
		mse := totalError / float64(len(data))
		fmt.Printf("Epoch %d: Weights: %v, Bias: %f, MSE: %f\n", epoch+1, layer.weights, layer.bias, mse)
	}
	return layer
}

// Evaluate the model
func predictAndEvaluate(data [][]float64, layer *Layer, mean []float64, std []float64) {
	features := len(data[0]) - 1
	totalError := 0.0
	for i := 0; i < len(data); i++ {
		standardizedFeatures := make([]float64, features)
		for j := 0; j < features; j++ {
			standardizedFeatures[j] = (data[i][j] - mean[j]) / std[j]
		}
		prediction := layer.Forward(standardizedFeatures)
		error := prediction - data[i][features]
		totalError += error * error
		fmt.Printf("Sample %d: Predicted Value: %f, Actual Value: %f\n", i+1, prediction, data[i][features])
	}
	mse := totalError / float64(len(data))
	fmt.Printf("Mean Squared Error: %f\n", mse)
}

func main() {
	// Read training data
	trainData, err := readCSV("train.csv")
	if err != nil {
		log.Fatalf("failed to read training data: %v", err)
	}

	// Read testing data
	testData, err := readCSV("test.csv")
	if err != nil {
		log.Fatalf("failed to read testing data: %v", err)
	}

	// Standardize training data
	standardizedTrainData, mean, std := standardize(trainData)

	// Train model
	learningRate := 0.01
	epochs := 1000
	layer := trainModel(standardizedTrainData, learningRate, epochs)

	// Evaluate model on test data
	predictAndEvaluate(testData, layer, mean, std)
}
