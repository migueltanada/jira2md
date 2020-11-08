Feature List

1. Supported Conversions
    
    ### ✓ Headings

    To create a header, place "hn. " at the start of the line (where n can be a number from 1-6).

    |   | Syntax               | Result in MD            |
    |---|----------------------|-------------------------|
    | ✓ | h1. Biggest heading  | # Biggest heading       |
    | ✓ | h2. Bigger heading   | ## Bigger heading       |
    | ✓ | h3. Big heading      | ### Big heading         |
    | ✓ | h4. Normal heading   | #### Normal heading     |
    | ✓ | h5. Small heading    | ##### Small heading     |
    | ✓ | h6. Smallest heading | ###### Smallest heading |

    ### Text Effects

    Text effects are used to change the formatting of words and sentences.

    |   | Syntax                                                                  | Result in MD                                                                                                |
    |---|-------------------------------------------------------------------------|-------------------------------------------------------------------------------------------------------------|
    | ✓ | *strong*                                                                | Makes text **strong**.                                                                                      |
    | ✓ | _emphasis_                                                              | Makes text *emphasis*..                                                                                     |
    | ✓ | ??citation??                                                            | Makes text in <cite>citation</cite>.                                                                        |
    |   | -deleted-                                                               | Makes text as ~~deleted~~.                                                                                  |
    | ✓ | +inserted+                                                              | Makes text as <ins>inserted</ins>.                                                                          |
    | ✓ | ^superscript^                                                           | Makes text in <sup>superscript</sup>.                                                                       |
    | ✓ | ~subscript~                                                             | Makes text in <sub>subscript</sub>.                                                                         |
    | ✓ | {{monospaced}}                                                          | Makes text as `monospaced`.                                                                                 |
    |   | bq. Some block quoted text                                              | To make an entire paragraph into a block quotation, place "bq. " before it. <br> > Example: Some block quoted text<br> |
    |   | {quote}<br>     here is quotable<br>  content to be quoted <br>{quote}  | Quote a block of text that's longer than one paragraph. Example:<br> ><br> > here is quotable <br> > content to be quoted <br> >      |
    |   | Markdown doesn't support color!<br>{color:red}<br> but html does! <br>{color}                      | Markdown doesn't support color!<br> <span style="color:red">but html does!</span>                               |

    ### Text Breaks

    Most of the time, explicit paragraph breaks are not required - The wiki renderer will be able to paginate your paragraphs properly.

    |   | Syntax       | Result in MD                                                                                                           |
    |---|--------------|------------------------------------------------------------------------------------------------------------------------|
    |   | (empty line) | Produces a new paragraph                                                                                               |
    |   | \\           | Creates a line break. Not often needed, most of the time the wiki renderer will guess new lines for you appropriately. |
    |   | ----         | Creates a horizontal ruler.                                                                                            |
    |   | ---          | Produces — symbol.                                                                                                     |
    |   | --           | Produces – symbol.                                                                                                     |

    ### Links

    Learning how to create links quickly is important.

    |   | Syntax                                                            | Result in MD                                                                                                                                                                                                                                                                                              |
    |---|-------------------------------------------------------------------|-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
    |   | [#anchor] [^attachment.ext]                                       | Creates an internal hyperlink to the specified anchor or attachment. Appending the '#' sign followed by an anchor name will lead into a specific bookmarked point of the desired page. Having the '^' followed by the name of an attachment will lead into a link to the attachment of the current issue. |
    |   | [http://jira.atlassian.com] [Atlassian\\\|http://atlassian.com]     | Creates a link to an external resource, special characters that come after the URL and are not part of it must be separated with a space. The [] around external links are optional in the case you do not want to use any alias for the link. Examples: http://jira.atlassian.com Atlassian              |
    |   | [mailto:legendaryservice@atlassian.com]                           | Creates a link to an email address, complete with mail icon. Example: #legendaryservice@atlassian.com                                                                                                                                                                                                     |
    |   | [file:///c:/temp/foo.txt] [file:///z:/file/on/network/share.txt]  | Creates a download link to a file on your computer or on a network share that you have mapped to a drive. To access the file, you must right click on the link and choose "Save Target As". By default, this only works on Internet Explorer but can also be enabled in Firefox (see docs).               |
    |   | {anchor:anchorname}                                               | Creates a bookmark anchor inside the page. You can then create links directly to that anchor. So the link [My Page#here] will link to wherever in "My Page" there is an {anchor:here} macro, and the link [#there] will link to wherever in the current page there is an {anchor:there} macro.            |
    |   | [~username]                                                       | Creates a link to the user profile page of a particular user, with a user icon and the user's full name.                                                                                                                                                                                                  |

    ### Lists

    Lists allow you to present information as a series of ordered items.

    |   | Syntax                                                             | Result in MD                                                                                                                                     |
    |---|--------------------------------------------------------------------|--------------------------------------------------------------------------------------------------------------------------------------------------|
    |   | * some <br>* bullet <br>** indented <br>** bullets <br>* points    | A bulleted list (must be in first column). Use more (**) for deeper indentations.<br>Example:<br>some<br>bullet<br>indented<br>bullets<br>points |
    |   | - different <br>- bullet <br>- types                               | A list item (with -), several lines create a single list.<br>Example:<br>different<br>bullet<br>types                                            |
    |   | # a <br># numbered<br><br># list                                   | A numbered list (must be in first column). Use more (##, ###) for deeper indentations.<br>Example:<br>a<br>numbered<br>list                      |
    |   | # a<br># numbered<br>#* with<br>#* nested<br>#* bullet<br># list   | You can even go with any kind of mixed nested lists<br>Example:<br>a<br>numbered<br>with<br>nested<br>bullet<br>list                             |
    |   | * a<br>* bulleted<br>*# with<br>*# nested<br>*# numbered<br>* list | Example:<br>a<br>bulleted<br>with<br>nested<br>numbered<br>list                                                                                  |

    ### Images

    Images can be embedded into a wiki renderable field from attached files or remote sources.

    |   | Syntax                                                   | Result in MD                                                                                                                                                            |
    |---|----------------------------------------------------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
    |   | !http://www.host.com/image.gif! or !attached-image.gif!  | Inserts an image into the page.<br>If a fully qualified URL is given the image will be displayed from the remote source, otherwise an attached image file is displayed. |
    |   | !image.jpg\|thumbnail!                                   | Insert a thumbnail of the image into the page (only works with images that are attached to the page).                                                                   |
    |   | !image.gif\|align=right, vspace=4!                       | For any image, you can also specify attributes of the image tag as a comma separated list of name=value pairs like so.                                                  |

    ### ✓ Tables

    Tables allow you to organise content in a rows and columns, with a header row if required.


    |   | Syntax                                                   | Result in MD       |
    |---|----------------------------------------------------------|--------------------|
    | ✓ | \|\| column1 \|\| column2 \|\| column3 \|\|  | \| column1 \| column2 \| column3 \|<br>\|-------\|-------\|-------\| |

    ### ✓ Advanced Formatting
    More advanced text formatting.

    | ✓ | Syntax                                                                                        | Result in MD                                                                                   |
    |---|-----------------------------------------------------------------------------------------------|------------------------------------------------------------------------------------------------|
    |   | {code:java}// Some comments here<br>public String getFoo()<br>{<br>    return foo;<br>}{code} | \`\`\`java<br>// Some comments here<br>public String getFoo()<br>{<br>    return foo;<br>}<br>\`\`\` |



 


