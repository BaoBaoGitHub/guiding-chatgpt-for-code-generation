# guiding-chatgpt-for-code-generation

The project include ChatGPT scripts, dataset and evaluators.

## How to use

### ChatGPT for Python

1. Install dependencies.    `pip install -r pydir/requirements.txt`
2. Set environment. Modify `OPENAI_API_KEY` in pydir/.env to your OpenAI API key.
3. Run script.    `python pydir/chat.py`

### ChatGPT for Golang

1. `cd godir`
2. Set PC environment variables.    `GO111MODULE=on`
3. Install dependencies.    `go mod download`
4. Run script.    `go run chat.go`

### Evaluator

1. Calculate BLEU.    `python evaluator/bleu/evaluator.py --ref "references path" --pre "predictions path"` 
2. Calculate CodeBLEU.
   1. `cd evaluator/codebleu`
   2. `python calc_code_bleu.py --refs "references path" --hyp "predictions path" --lang java`

## Dataset

1. dataset/text2code includes dataset of text to code generation.
2. dataset/code2code includes dataset of code to code translation.