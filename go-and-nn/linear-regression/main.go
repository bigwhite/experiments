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

// Helper function to train the model
func trainModel(data [][]float64, learningRate float64, epochs int) ([]float64, float64) {
	features := len(data[0]) - 1
	weights := make([]float64, features)
	bias := 0.0

	for epoch := 0; epoch < epochs; epoch++ {
		gradW := make([]float64, features)
		gradB := 0.0
		mse := 0.0
		for i := 0; i < len(data); i++ {
			prediction := bias
			for j := 0; j < features; j++ {
				prediction += weights[j] * data[i][j]
			}
			error := prediction - data[i][features]
			mse += error * error
			for j := 0; j < features; j++ {
				gradW[j] += error * data[i][j]
			}
			gradB += error
		}
		mse /= float64(len(data))

		for j := 0; j < features; j++ {
			gradW[j] /= float64(len(data))
			weights[j] -= learningRate * gradW[j]
		}
		gradB /= float64(len(data))
		bias -= learningRate * gradB

		// Output the current weights, bias and loss
		fmt.Printf("Epoch %d: Weights: %v, Bias: %f, MSE: %f\n", epoch+1, weights, bias, mse)
	}
	return weights, bias
}

// Helper function to predict and evaluate the model
func predictAndEvaluate(data [][]float64, weights []float64, bias float64, mean []float64, std []float64) {
	features := len(data[0]) - 1
	mse := 0.0
	for i := 0; i < len(data); i++ {
		// Standardize the input features using the training mean and std
		standardizedFeatures := make([]float64, features)
		for j := 0; j < features; j++ {
			standardizedFeatures[j] = (data[i][j] - mean[j]) / std[j]
		}

		// Calculate the prediction
		prediction := bias
		for j := 0; j < features; j++ {
			prediction += weights[j] * standardizedFeatures[j]
		}

		// Calculate the error and accumulate the MSE
		error := prediction - data[i][features]
		mse += error * error

		// Print the prediction and the actual value
		fmt.Printf("Sample %d: Predicted Value: %f, Actual Value: %f\n", i+1, prediction, data[i][features])
	}

	// Calculate the final MSE
	mse /= float64(len(data))
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

	// Plot the dot
	//drawDots(testData, 0, "area", "price")
	//drawDots(testData, 1, "distance", "price")

	// Standardize training data
	standardizedTrainData, mean, std := standardize(trainData)

	// Train model
	//learningRate := 0.0001
	learningRate := 0.01
	epochs := 1000
	weights, bias := trainModel(standardizedTrainData, learningRate, epochs)
	fmt.Printf("Trained Weights: %v\n", weights)
	fmt.Printf("Trained Bias: %f\n", bias)

	// Evaluate model on test data
	predictAndEvaluate(testData, weights, bias, mean, std)
}
