// matrix.cu
#include <cuda_runtime.h>
#include <stdio.h>

// CUDA Kernel: 每个线程计算 C[row][col] 的值
__global__ void matrixMulKernel(float *a, float *b, float *c, int width) {
    // 根据 Block ID 和 Thread ID 计算当前线程的全局坐标
    int row = blockIdx.y * blockDim.y + threadIdx.y;
    int col = blockIdx.x * blockDim.x + threadIdx.x;

    if (row < width && col < width) {
        float sum = 0;
        // 计算点积
        for (int k = 0; k < width; k++) {
            sum += a[row * width + k] * b[k * width + col];
        }
        c[row * width + col] = sum;
    }
}

extern "C" {
    // 供 Go 调用的 C 包装函数
    // 负责显存分配、数据拷贝和内核启动
    void runMatrixMul(float *h_a, float *h_b, float *h_c, int width) {
        int size = width * width * sizeof(float);
        float *d_a, *d_b, *d_c;

        // 1. 分配 GPU 显存 (Device Memory)
        cudaMalloc((void **)&d_a, size);
        cudaMalloc((void **)&d_b, size);
        cudaMalloc((void **)&d_c, size);

        // 2. 将数据从 Host (CPU内存) 复制到 Device (GPU显存)
        // 这一步通常是性能瓶颈，应尽量减少
        cudaMemcpy(d_a, h_a, size, cudaMemcpyHostToDevice);
        cudaMemcpy(d_b, h_b, size, cudaMemcpyHostToDevice);

        // 3. 定义 Grid 和 Block 维度
        // 每个 Block 包含 16x16 = 256 个线程
        dim3 threadsPerBlock(16, 16);
        // Grid 包含足够多的 Block 以覆盖整个矩阵
        dim3 numBlocks((width + threadsPerBlock.x - 1) / threadsPerBlock.x,
                       (width + threadsPerBlock.y - 1) / threadsPerBlock.y);

        // 4. 启动内核！成千上万个线程开始并行计算
        matrixMulKernel<<<numBlocks, threadsPerBlock>>>(d_a, d_b, d_c, width);

        // 5. 将计算结果从 Device 传回 Host
        cudaMemcpy(h_c, d_c, size, cudaMemcpyDeviceToHost);

        // 6. 释放 GPU 内存
        cudaFree(d_a);
        cudaFree(d_b);
        cudaFree(d_c);
    }
}
