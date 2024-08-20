# ASCII-ART-WEB PROGRAM

Ascii-art-web is an improved implementation of the Ascii-art program.Ascii-art-web consists in creating and running a server, in which it will be possible to use a web GUI (graphical user interface) version to input and display the art in ascii format.
Ascii-art is a program which consists in receiving a string as an argument and outputting the string in a graphic representation using ASCII. Time to write big.

The program receives a string as an argument and outputs the string in a graphic representation using ASCII. All this is done on a webpage.

## Authors
- [Valeria Muhembele](https://github.com/anamivale)
- [Sheilla Juma](https://github.com/a-j-sheilla)


## Usage: How to run
To run the code, in the terminal, run

  ```sh
  go run main.go
  ```
  This will prompt the program to run and execute on http://localhost:1024
  Open the web browser and in the search bar, search and load [localhost:1024].

You will get be provided a user interface where you will input your text, select an ascii art format of your choice among the three available options([standard], [thinkertoy] and [shadow]).
After selecting the desired ascii art format, click generate to get your output.

## Implementation Details: Algorithm

The ASCII Art Web application is designed to render text in different ASCII art styles (banners) based on user input. Below is a detailed description of the algorithm and logic used to implement the project:

### 1. Banner Selection

The application supports three ASCII art styles (banners):
- Shadow
- Standard
- Thinkertoy

Each style is represented by a different font file containing the ASCII art representations of characters. The files are read and processed to generate the art.

### 2. Processing the Input

1. **Input Validation**: 
   - Validate user input to ensure it only contains supported characters. This prevents errors in rendering ASCII art.

2. **Text to ASCII Conversion**:
   - For each character in the input text:
     - Look up the ASCII art representation from the selected banner file.
     - Append the corresponding ASCII art lines to build the complete text representation.
   
3. **Line-by-Line Assembly**:
   - ASCII art is generated line by line, respecting the height of the characters as defined by the banner.
   - For example, if each character's representation is 8 lines tall, the entire text is assembled by concatenating corresponding lines from each character's representation.

### 3. HTTP Server Logic

#### Endpoints

- **GET /**:
  - Renders the main HTML page with a form for text input, radio buttons (or a select menu) to choose the banner style, and a submit button.
  - Utilizes Go templates to dynamically render content.

- **POST /**:
  - Receives form data from the client containing the text and selected banner style.
  - Processes the input to generate the ASCII art using the logic described above.
  - Sends back the ASCII art as part of the HTTP response.

#### Error Handling

- **200 OK**: Returned when the request is processed successfully.
- **404 Not Found**: Returned when a requested resource (e.g., a template or banner file) cannot be found.
- **400 Bad Request**: Returned for invalid inputs or requests that do not conform to expected formats.
- **500 Internal Server Error**: Returned for unhandled exceptions or errors during processing.

### 4. User Interface

- **HTML Form**: Allows users to enter text and select the banner style.
- **JavaScript (optional)**: Can be used to enhance user interaction, such as displaying a loading indicator during processing.

The combination of server-side processing and a user-friendly web interface allows for an intuitive and efficient ASCII art generation experience.
