# 📦 Ascii-Art-Web-Export

Ascii-Art-Web-Export is an extension of the Ascii Art Web project, providing the capability to export the generated ASCII art to a file format (.txt). This feature ensures that users can download their ASCII art creations directly from the web interface, with proper file permissions set for reading and writing.

## 📄 Description

Ascii-Art-Web-Export enhances the original ASCII Art Web project by allowing users to export their generated ASCII art as a downloadable file. This extension includes the implementation of HTTP headers such as `Content-Type`, `Content-Length`, and `Content-Disposition` to ensure the file is transferred correctly to the client.

### ✨ Features

- **Export ASCII Art:** Allows users to export the generated ASCII art in a chosen file format.
- **HTTP Header Implementation:** Ensures the correct use of HTTP headers for file transfer.
- **Download Button:** A convenient button or link on the web interface for easy downloading.
- **Error Handling:** Manages errors related to file generation and transfer gracefully.

## 🛠 How It Works

1. **Generate ASCII Art:** The web application continues to allow users to input text and select a banner style to generate ASCII art.
2. **Export to File:** Once the ASCII art is generated, users can click the download button to export the ASCII art.
3. **HTTP Headers:** The server uses specific HTTP headers to manage the file download process.
4. **File Permissions:** The exported file is created with the appropriate read and write permissions for the user.

## 🔍 Implementation

- **Go:** Handles backend logic, HTTP requests, template rendering, and file export functionality.
- **HTML:** Structures the web interface, including the export/download button.
- **CSS:** Styles the interface, ensuring a consistent and user-friendly design.

### Algorithms

- **HTTP Server:** The Go server handles the creation of a new HTTP endpoint specifically for file export.
- **Header Management:** The program sets the required HTTP headers (`Content-Type`, `Content-Length`, `Content-Disposition`) to manage file transfer.
- **Error Handling:** The server checks for issues like missing ASCII art or file generation errors and responds accordingly.

## 🎨 HTML and CSS Usage

### HTML
- The HTML file is structured to include a user-friendly interface with an additional download button for exporting ASCII art.

### CSS
- The CSS file styles the HTML elements, ensuring the download button is visually integrated into the existing design.

## 💾 How to Install

Ascii-Art-Web-Export requires Go v1.22.2 to run.

**Clone the repository:**

    ```bash
    git clone https://github.com/krodgers/ascii-art-web-export.git
    cd ascii-art-web-export
    ```

## ▶️ Running the Program

To run the program, use the following command:

```bash
go run .
```
The terminal will prompt:

```
Starting server at :8080

To access the web interface, open your web browser and navigate to http://localhost:8080.
```

## 💡 Usage
Access the Web Interface:

Open your web browser and navigate to the provided URL (i.e., http://localhost:8080).
Generate and Export ASCII Art:

   1. Input Text: Enter the text you want to convert into ASCII art.
   2. Select a Banner Style: Choose from available styles like Standard, Thinkertoy, or Shadow.
   3. Generate ASCII Art: Click the "Generate ASCII Art" button to see the output.
   4. Export ASCII Art: Click the "Download ASCII Art" button to save the generated ASCII art as a file to your device.

Reset Input:

Use the "Reset" button to clear the input and start over.
Interact with the Output:

Scroll within the ASCII art box to view larger outputs and adjust the view as needed.


## ✍️ Authors

<table>
<tr>
    <td align="center" style="word-wrap: break-word; width: 150.0; height: 150.0">
        <a href=https://www.linkedin.com/in/rodgers-kaunda>
            <img src=https://learn.zone01kisumu.ke/git/avatars/aa19095145ab1ad43695e3cd3f7f3a5b?size=870 width="100;"  style="border-radius:50%;align-items:center;justify-content:center;overflow:hidden;padding-top:10px" alt=krodgers/>
            <br />
            <sub style="font-size:14px"><b>krodgers</b></sub>
        </a>
    </td>
    <td align="center" style="word-wrap: break-word; width: 150.0; height: 150.0">
        <a href=https://www.linkedin.com/in/swabri-musa-565350291?lipi=urn%3Ali%3Apage%3Ad_flagship3_profile_view_base_contact_details%3Buf0Ls4oWR2O2WLUMO5sIBg%3D%3D>
            <img src=https://learn.zone01kisumu.ke/git/avatars/bc7899a0aac2630a0a9b50bf330437a7?size=870 width="100;"  style="border-radius:50%;align-items:center;justify-content:center;overflow:hidden;padding-top:10px" alt=skanenje/>
            <br />
            <sub style="font-size:14px"><b>skanenje</b></sub>
        </a>
    </td>

</tr>
</table>

Feel free to contribute to this project by opening issues or submitting pull requests!!