package test

import (
	"testing"

	"github.com/nardineshak/fast-tiktoken/codec"
)

var sentences []string = []string{"This is a long sentence that includes special characters like @, #, $ and % to test the OpenAI's tokenizer's ability to identify and separate these symbols.",
	"Sometimes, sentences use numbers like 123, 456.78, and +9, -7 or 8% to describe quantities, amounts, percentages, and it's important for a tokenizer to recognize these.",
	"Tokenizing code : for i in range(10) { fmt.Println(i) } can be tricky as they include parentheses, braces, colons, loops, print statements and more.",
	"While the world witnessed a paradigm shift in technologies with the advent of advanced Machine Learning, NLP, Deep Learning and their convergence with Cloud Computing; it has become essential for scientists to adapt to these changes rapidly.",
	"Quick brown fox jumps over the lazy dog, isn't that what we always say when testing typewriters or any tool that involves the use of alphabetical characters?",
	"Can AI comprehend symbolic relations such as if X is greater than Y and Z is less than or equal to Y, then definitely X is greater than Z?",
	"Go consider a != 0 && (1/a > 1/b) in C++, Python, JavaScript, Perl and if it computes rightly, then your tokenizer and the subsequent parser are doing a commendable job!",
	"It was Capt. James who said, 'I have 3 missions: 1. To explore brave new worlds. 2. To seek out new life. 3. To boldly go where no man has gone before.'",
	"The Emancipation Proclamation in 1863 was a significant chapter in American history; it changed the course of the Civil War, precipitated the passage of the 13th Amendment and marked the beginning of the end for slavery.",
	"Containing an extensive range of punctuation marks - like hyphens, ellipsis, brackets (both round and square) - can pose considerable challenges to tokenizers.",
	"Tokenizing sentences like 'The COVID-19 pandemic drastically changed the world, leading to unprecedented shifts in work paradigms, pushing global economies, and disrupting traditional societal structures' can push a tokenizer to its full potential.",
	"How could tokenizers like OpenAI's handle Incorporated (Inc.), Limited (Ltd.), Doctor (Dr.), Mister (Mr.), and characters encased within brackets followed by periods?",
	"Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.",
	"Tokenize this, master-commander: docker pull ubuntu:latest || docker run -it ubuntu:latest /bin/bash - quite intriguing with Linux commands, isn’t it?",
	"Can a tokenizer distinguish between 'It's 8 P.M. now' and 'St. Patrick's Day is on Mar. 17 this year' and create an array of tokens aptly?",
	"Sentences with email addresses like john.doe@example.com and URLs like https://www.example.com can be tricky for tokenizers, can they deal with these?",
	"She expressed, 'Je suis heureuse', which translated to 'I am happy' in English; it's quite interesting to see how different languages get tokenized.",
	"The classic opening line 'It was the best of times, it was the worst of times...' from Charles Dickens' 'A Tale of Two Cities' may seem simple, but accurate tokenization of literary text can be quite the task.",
	"It's fascinating to consider how a tokenizer handles alphanumeric combinations, like 'My flight number is SQ346, it departs at 23:45 on 3rd Jan 2022 from Terminal 3B.'",
	"'ChVrches' and 'The Wknd' are both popular music bands, but does the tokenizer capture the uniqueness in the spelling of their names?"}

var paragraph1 string = "Natural language processing (NLP) is a field of artificial intelligence (AI) that focuses on enabling computers to understand, interpret, and generate human language in a way that is both meaningful and useful. It involves a range of tasks such as text tokenization, part-of-speech tagging, named entity recognition, sentiment analysis, and more. NLP has numerous applications in various industries including healthcare, finance, customer service, and language translation. Developing efficient tokenization methods is crucial for many NLP tasks as it forms the basis for further analysis and processing of textual data."

var paragraph2 string = "OpenAI's large language models (sometimes referred to as GPT's) process text using tokens, which are common sequences of characters found in a set of text. The models learn to understand the statistical relationships between these tokens, and excel at producing the next token in a sequence of tokens.You can use the tool below to understand how a piece of text might be tokenized by a language model, and the total count of tokens in that piece of text.It's important to note that the exact tokenization process varies between models. Newer models like GPT-3.5 and GPT-4 use a different tokenizer than our legacy GPT-3 and Codex models, and will produce different tokens for the same input text."

func benchmarkTokenize(str string, b *testing.B) {
	for i := 0; i < b.N; i++ {
		codec.Tokenize(str)
	}
}

func benchmarkEstimateTokens(str string, b *testing.B) {
	for i := 0; i < b.N; i++ {
		codec.EstimateTokens(str)
	}
}

func BenchmarkTokenizeSentences(b *testing.B) {
	for _, str := range sentences {
		benchmarkTokenize(str, b)
	}
}

func BenchmarkEstimateTokensSentences(b *testing.B) {
	for _, str := range sentences {
		benchmarkEstimateTokens(str, b)
	}
}

func BenchmarkTokenizeParagraph1(b *testing.B)       { benchmarkTokenize(paragraph1, b) }
func BenchmarkEstimateTokensParagraph1(b *testing.B) { benchmarkEstimateTokens(paragraph1, b) }

func BenchmarkTokenizeParagraph2(b *testing.B)       { benchmarkTokenize(paragraph2, b) }
func BenchmarkEstimateTokensParagraph2(b *testing.B) { benchmarkEstimateTokens(paragraph2, b) }

// func BenchmarkTokenizeS1(b *testing.B) { benchmarkTokenize(sentences[0], b) }
// func BenchmarkTokenizeS2(b *testing.B) { benchmarkTokenize(sentences[1], b) }
// func BenchmarkTokenizeS3(b *testing.B) { benchmarkTokenize(sentences[2], b) }

// func BenchmarkEstimateTokensS1(b *testing.B) { benchmarkEstimateTokens(sentences[0], b) }
// func BenchmarkEstimateTokensS2(b *testing.B) { benchmarkEstimateTokens(sentences[1], b) }
// func BenchmarkEstimateTokensS3(b *testing.B) { benchmarkEstimateTokens(sentences[2], b) }
