package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// Activation function (Sigmoid)
func sigmoid(x float64) float64 {
	return 1.0 / (1.0 + math.Exp(-x))
}

// Derivative of the sigmoid function
func sigmoidDerivative(x float64) float64 {
	return x * (1.0 - x)
}

// MLP structure
type MLP struct {
	inputLayer          []float64
	hiddenLayer         []float64
	outputLayer         []float64
	weightsInputHidden  [][]float64
	weightsHiddenOutput []float64
	learningRate        float64
}

// Initialize the MLP
func (mlp *MLP) Initialize(inputSize, hiddenSize, outputSize int, learningRate float64) {
	mlp.inputLayer = make([]float64, inputSize)
	mlp.hiddenLayer = make([]float64, hiddenSize)
	mlp.outputLayer = make([]float64, outputSize)
	mlp.weightsInputHidden = make([][]float64, inputSize)
	for i := 0; i < inputSize; i++ {
		mlp.weightsInputHidden[i] = make([]float64, hiddenSize)
		for j := 0; j < hiddenSize; j++ {
			mlp.weightsInputHidden[i][j] = randWeight()
		}
	}
	mlp.weightsHiddenOutput = make([]float64, hiddenSize)
	for i := 0; i < hiddenSize; i++ {
		mlp.weightsHiddenOutput[i] = randWeight()
	}
	mlp.learningRate = learningRate
}

// Forward pass
func (mlp *MLP) Forward(inputs []float64) []float64 {
	// Input to Hidden
	for j := 0; j < len(mlp.hiddenLayer); j++ {
		mlp.hiddenLayer[j] = 0
		for i := 0; i < len(mlp.inputLayer); i++ {
			mlp.hiddenLayer[j] += inputs[i] * mlp.weightsInputHidden[i][j]
		}
		mlp.hiddenLayer[j] = sigmoid(mlp.hiddenLayer[j])
	}

	// Hidden to Output
	for k := 0; k < len(mlp.outputLayer); k++ {
		mlp.outputLayer[k] = 0
		for j := 0; j < len(mlp.hiddenLayer); j++ {
			mlp.outputLayer[k] += mlp.hiddenLayer[j] * mlp.weightsHiddenOutput[j]
		}
		mlp.outputLayer[k] = sigmoid(mlp.outputLayer[k])
	}

	return mlp.outputLayer
}

// Training using backpropagation
func (mlp *MLP) Train(inputs [][]float64, targets [][]float64, epochs int) {
	for epoch := 0; epoch < epochs; epoch++ {
		for idx, input := range inputs {
			outputs := mlp.Forward(input)

			// Calculate output layer errors and deltas
			outputErrors := make([]float64, len(mlp.outputLayer))
			outputDeltas := make([]float64, len(mlp.outputLayer))
			for k := 0; k < len(mlp.outputLayer); k++ {
				outputErrors[k] = targets[idx][k] - outputs[k]
				outputDeltas[k] = outputErrors[k] * sigmoidDerivative(outputs[k])
			}

			// Calculate hidden layer errors and deltas
			hiddenErrors := make([]float64, len(mlp.hiddenLayer))
			hiddenDeltas := make([]float64, len(mlp.hiddenLayer))
			for j := 0; j < len(mlp.hiddenLayer); j++ {
				hiddenErrors[j] = 0
				for k := 0; k < len(mlp.outputLayer); k++ {
					hiddenErrors[j] += outputDeltas[k] * mlp.weightsHiddenOutput[j]
				}
				hiddenDeltas[j] = hiddenErrors[j] * sigmoidDerivative(mlp.hiddenLayer[j])
			}

			// Update weights for Hidden to Output
			for j := 0; j < len(mlp.hiddenLayer); j++ {
				for k := 0; k < len(mlp.outputLayer); k++ {
					mlp.weightsHiddenOutput[j] += mlp.learningRate * outputDeltas[k] * mlp.hiddenLayer[j]
				}
			}

			// Update weights for Input to Hidden
			for i := 0; i < len(mlp.inputLayer); i++ {
				for j := 0; j < len(mlp.hiddenLayer); j++ {
					mlp.weightsInputHidden[i][j] += mlp.learningRate * hiddenDeltas[j] * input[i]
				}
			}
		}

		if epoch%1000 == 0 {
			error := 0.0
			for i, input := range inputs {
				outputs := mlp.Forward(input)
				for k := 0; k < len(mlp.outputLayer); k++ {
					error += math.Pow(targets[i][k]-outputs[k], 2)
				}
			}
			fmt.Printf("Epoch %d, Error: %f\n", epoch, error)
		}
	}
}

// Helper function to generate random weight
func randWeight() float64 {
	return rand.Float64()*2 - 1 // Random weight between -1 and 1
}

// Main function
func main() {
	rand.Seed(time.Now().UnixNano())

	inputs := [][]float64{
		{0, 0},
		{0, 1},
		{1, 0},
		{1, 1},
	}

	targets := [][]float64{
		{0},
		{1},
		{1},
		{0},
	}

	mlp := MLP{}
	mlp.Initialize(2, 2, 1, 0.1) // Increased hidden layer size to 2

	mlp.Train(inputs, targets, 20000) // Increased epochs to 20000

	fmt.Println("Trained model parameters:")
	fmt.Println("Hidden Layer Weights:", mlp.weightsInputHidden)
	fmt.Println("Output Layer Weights:", mlp.weightsHiddenOutput)

	fmt.Println("\nTesting the neural network:")
	for _, input := range inputs {
		predicted := mlp.Forward(input)
		class := 0
		if predicted[0] >= 0.5 {
			class = 1
		}
		fmt.Printf("Input: %v, Predicted: %v, Classified as: %d, Actual: %v\n", input, predicted, class, targets)
	}
}
