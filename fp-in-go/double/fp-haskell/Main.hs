import System.IO
import Control.Monad (when)
import Text.Read (readMaybe)
import Data.Maybe (catMaybes)

-- Define a custom type for the result
data DoubledNumbers = DoubledNumbers { doubledNumbers :: [Int] } deriving (Show)

-- Function to read numbers from a file
readNumbers :: FilePath -> IO (Either String [Int])
readNumbers filePath = do
    content <- readFile filePath
    let numbers = catMaybes (map readMaybe (lines content))
    return $ if null numbers
             then Left "No valid numbers found."
             else Right numbers

-- Function to write result to a file
writeResult :: FilePath -> DoubledNumbers -> IO (Either String ())
writeResult filePath result = do
    let resultString = unlines (map show (doubledNumbers result))
    writeFile filePath resultString
    return $ Right ()

-- Function to double the numbers
doubleNumbers :: [Int] -> DoubledNumbers
doubleNumbers numbers = DoubledNumbers { doubledNumbers = map (* 2) numbers }

main :: IO ()
main = do
    -- Read numbers from input.txt
    readResult <- readNumbers "input.txt"
    case readResult of
        Left err -> putStrLn $ "Error: " ++ err
        Right numbers -> do
            let result = doubleNumbers numbers
            -- Write result to output.txt
            writeResultResult <- writeResult "output.txt" result
            case writeResultResult of
                Left err -> putStrLn $ "Error: " ++ err
                Right () -> putStrLn "Successfully written the result to output.txt."
