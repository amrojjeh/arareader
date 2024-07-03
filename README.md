بسم الله الرحمن الرحيم

An open-source curriculum for learning Arabic, with a focus on guided readings and exercises.

# Technical terms
- *Guided Reading*: A collection of questions and a single excerpt
- *Question*: At minimum, a question must have a prompt and some method of inputting the answer
- *Excerpt*: An excerpt is the collection of Arabic words which is subject to highlighting
- *Short vowel*: The dammah, fatha, kasra, their doubled equivalents, and the sukoon
- *Immediate mode*: Answers can be submitted as soon as they are attempted
- *Ultimate mode*: Answers are submitted at the very end after all questions have been attempted (should probably change the name)

## Re: Excerpts
> **Needs to be updated**

Excerpts will be stored in a plaintext format. Here's an example:

```
هذا بي١تٌ وـ٢ذلك شي٣ءٌ٢ـ.
0: SHORT ANSWER
Prompt: Translate
Answer: This is a house and that is a thing
1: VOWEL
2: SHORT ANSWER
Prompt: Translate
Answer: and that is a thing
Grade: self
3: VOWEL
```

Note that the references begin counting from 1, since the entire excerpt would be reference zero. Also note that references can be inside of each other.

In this example, the first question is short answer question (highlights nothing), the second is a vowel question (highlights تٌ), the third is short answer question (highlights ذلك شيءٌ), and the last question is another vowel question (highlights ءٌ).

The `Grade: self` means that rather than strictly comparing the student's answer with the provided one, the app will give the user the answer and they'll decide if it's close enough. In the future, we could possibly implement a mentoring system where someone else would check the translation, but I believe students will be capable of grading themselves in most questions.

References can be reused for multiple questions.

## Re: Questions
- *Question type*: Differentiated by these qualities:
  - Whole/Letter Segmented/Word Segmented/Phrase Segmented
  - Method of input
  - Method of validation (self or automatic)
- *Whole*: A question which does not highlight any part of the excerpt (reference 0)
- *Segment question*: A question which includes a reference to a substring in the excerpt
- *Letter Segmented*: Segmented question which only highlights a single letter (including its diacritical marks)
- *Phrase Segmented*: Segmented question which only highlights a collection of words (could just be one word)

In the excerpt, the *letter segment* is indicated by just placing right before the letter with no spacing (`٣دعاء` highlights the د). The *phrase segment* is indicated by placing a tatweel before the number, and also ending it with the same number and another tatweel (`ـ٣دعاء٣ـ` highlights the entire word).

## Re: Question types
- *Short answer question*:
  - Letter/Phrase Segmented
  - Student types in answer. Restrictions:
    - 40 Character limit
  - Self/Automatic
- *Vowel question*:
  - Letter Segmented
  - Student chooses a short vowel
  - Automatic


### Potential question types
- *Long answer question*:
  - Letter/Phrase segmented
  - Student types in answer
    - 200 character limit
  - Self/Automatic
- *Word pronunciation question*
  - Letter/Phrase Segmented
  - Student records pronunciation
  - Self
- *Substitution question*
  - Letter/Phrase Segmented
  - Student either selects or types in the correct word
  - Automatic
- *PGN question*:
  - Phrase Segmented
  - Student picks person, gender, number. Supports PGN notation for shortcuts
  - Automatic
- *Number question*:
  - Letter/Phrase Segmented
  - Student types in any number
  - Automatic
- *Multiple choice question*:
  - Letter/Phrase Segmented
  - Student chooses an answer from a set predetermined choices
  - Automatic

Scrapped:

- *Form question*: Student picks the correct form of the verb. This is really just a number question except there's a constraint on the number... not that helpful
- *Translation question*: It's either a short answer or a long answer question.
