Like Google Forms (quiz mode) but specialized for long Arabic excerpts. Lets you highlight the relevant text for each question and it has built in support for testing vowels, choosing the correct word, short answers, and so on. Automatically grades what it can and leaves the rest to the teacher.

# Technical terms
- *Class*: A collection of students, quizzes and a single teacher
- *Quiz*: A collection of questions and a single excerpt
- *Question*: At minimum, a question must have a prompt, some method of input which is called the answer, and the answer must either be correct or incorrect.
- *Excerpt*: An excerpt is the collection of Arabic words which is subject to highlighting
  - In the future, an excerpt could also be in the data format of [[Text Annotator]]
- *Short vowel*: The dammah, fatha, kasra, their doubled equivalents, and the sukoon
- *Immediate mode*: Answers can be submitted as soon as they are attempted
- *Ultimate mode*: Answers are submitted at the very end after all questions have been attempted

## Re: Excerpts
Excerpts will be stored in the XML format. Here's an example (the Arabic is encoded in Buckwalter just for this example, but in actuality it would be in unicode):
```xml
<excerpt>h*A by<ref id="1">tN</ref> <ref id="2">wh*A $y<ref id="3">'N</ref></ref></excerpt>
```

Note that the references begin counting from 1, since the entire excerpt would technically be reference zero. Also note that references can be inside of each other.

In this example, the first reference could be a vowel question, the second reference a translation (short answer) question and the last is another vowel question.

## Re: Questions
- *Question type*: Differentiated by these qualities:
  - Whole/Letter Segmented/Word Segmented/Phrase Segmented
  - Method of input, including the domain of input
  - Method of validation
- *Whole*: A question which does not highlight any part of the excerpt
- *Segment question*: A question which includes a reference to a substring in the excerpt
- *Letter Segmented*: Segmented question which only highlights a single letter
- *Word Segmented*: Segmented question which only highlights a single word, without any punctuation
- *Phrase Segmented*: Segmented question which only highlights a collection of words (could just be one word)

## Re: Question types
- *Short answer question*:
  - Whole/Phrase Segmented
  - Student types in answer. Restrictions:
    - 140 Character limit
  - Teacher manually grades each answer
- *Long answer question*:
  - Whole/Phrase segmented
  - Student types in answer, no character restriction
  - Teacher manually grades each answer
- *Multiple choice question*:
  - Whole/Phrase Segmented
  - Student chooses an answer from a set determined by the teacher. Restrictions:
    - The set must be between two and six (inclusive)
  - Teacher predetermines answer, so student choices are graded automatically
- *Number question*:
  - Whole/Phrase Segmented
  - Student types in any number
  - Correct answer is predetermined, so answers are graded automatically
- *Vowel question*:
  - Letter Segmented
  - Student chooses a short vowel
  - Graded automatically based on the vowel in the excerpt
- *PGN question*:
  - Word Segmented (specifically limited to verbs)
  - Student picks person, gender, number. Supports PGN notation for shortcuts
  - Correct answer is predetermined by teacher, so answers are graded automatically

### Potential question types
- *Word pronunciation question*
  - Whole/Phrase Segmented
  - Student records pronunciation
  - Graded manually by teacher
- *Short Answer question*
  - Same as before but also support automatic grading using predetermined set of answers
  - Would be useful for a question such as "what is the root of this word?"
- *Substitution question*
	- Word Segmented
	- Student either selects or types in the correct word
	- Graded automatically
	- Useful if the teacher wants to test the student on verb form, plurality, gender, etc...

Scrapped:
- *Form question*: Student picks the correct form of the verb. This is really just a number question except there's a constraint on the number... not that helpful
- *Translation question*: It's either a short answer or a long answer question.

# FAQ

> [!question] Will it be free or commercial?
> Inshallah we can turn it into a commercial product, but for now I'm aiming to release it for free to at least gain an audience.

> [!question] Does it have a finished state or will it keep updating?
> It'll likely keep updating since it's a [[Services (Software)|service]].

> [!question] Why does Google Forms not work?
> Google forms cannot display Arabic properly. It cannot highlight relevant text. It's incapable of doing question types such as PGN or vowels.

> [!question] Would this be useful even in in-person environments?
> Yes. It saves paper and it allows students to get instantanious feedback. In the future, teachers may also have access to a plethora of resources which they simply add to their classes.

> [!question] Will it have built-in support for irab/grammatical annotations?
> Yes. This could be done using [[Text Annotator]]

> [!question] Should quizzes support multiple excerpts?
> Not at the moment. If it was to be supported, then excerpt questions should not be mixed together.
> The student should have to complete all questions related to an excerpt before moving on to the next one.

> [!question] Should classes support multiple teachers?
> No. For now, they can just share accounts. There does not seem to be a good reason to build this functionality in.

> [!question] When should student answers be evaluated?
> There are two choices:
> 1. Evaluate answers as soon as the student submits his answer to a question
> 2. Evaluate answers after the student completes all questions
>
> The first option limits space for review, but could help the student in upcoming questions. It also feels less dreadful to take. The second option offers space for review, but delays feedback until the very end, which is closer to reality and is likely more effective. Both options can be supported by placing a submit button on every question in the first case and placing only a single submit button at the end of the quiz in the second case.

> [!question] Ability to take quizzes without being in a class?
> One of the things that make Google Forms useful is that you can just share a link and the student is immediately ready to take the quiz. With the current system, while the student does not need to make an account, he must first be added by the teacher to a class, then he can take the quiz. However, since classes are not integral to quizzes, we can perhaps support both class management and link sharing, the former allowing for more features such as exam taking, limiting attempts, etc...


# Marketing
Invite only?

# Technical Tasks
First phase is development prior to teacher feedback/marketing
Second phase is development after feedback/marketing

## First Phase
- [x] Excerpt API
- [x] **Redesign figma to improve feedback location + phone support**
- [ ] Add logo
- [ ] Improve design (a broken window -- which we cannot leave behind)
  - [ ] Write a components page
  - [ ] Improve layout
- [ ] Write test cases for short answer
  - [ ] Add a character limit
- [x] Add support for quiz sessions
  - [x] Add submit button
  - [x] Add next/prev question
  - [x] Load correct question prompt
  - [x] Load submitted answers
  - [x] Disable question if already submitted
  - [x] Add testing
  - [ ] Support shortcuts for next, submit, plus nav
  - [x] Display feedback
- [ ] Student perspective
  - [ ] Add student home page
  - [ ] Add side drawer
  - [ ] Add quiz report/summary page
  - [x] Support vowel type question
  - [x] Support short answer question

## Second phase
- [ ] Use auth solutions like Clerk or OAuth
- [ ] A way to retake the quiz
- [ ] Add individualized feedback
- [ ] Student perspective
  - [ ] Add sign in
  - [ ] Support long answer question
  - [ ] Support multiple choice question
  - [ ] Number question
  - [ ] PGN question
- [ ] Teacher perspective
  - [ ] Class management
    - [ ] Adding/Removing classes
    - [ ] Adding/Removing quizzes
    - [ ] Adding/Removing students
    - [ ] Grading student submissions
    - [ ] Viewing student submissions
  - [ ] Quiz management
    - [ ] Write excerpt
      - [ ] Disallow English characters
      - [ ] Show references
      - [ ] Maintain stable references
  - [ ] Support vowel type question
  - [ ] Support short answer question
  - [ ] Support long answer question
  - [ ] Support multiple choice question
  - [ ] Number question
  - [ ] PGN question

## Later
- [ ] Later
  - [ ] Support [[Text Annotator]]
