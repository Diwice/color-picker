# Color Picker

Unfinished (and mostly likely will never be) - the more I dig into fyne, the more I need to do things with fyne - the worse and more complicated everything gets, and, currently, I'm at the point when I myself barely understand the structure of my own app. Decided to leave it be like that since I can't quite fix it (well, I could, but it'd take even MORE time with ). Could've finished the cli tool instead within a week, but finished up spending almost 2 weeks on a project with more and more problems popping up.

So, what exactly is unfinished? Here : 
1. Can't really add a separator not breaking grid layout for hex part (wanted to add a thin line between color display & block of hex entry/check)
2. Hex part data isn't binded (hex entry, color display rectangle color)
3. Check should affect text color according to data display's color (some kind of shift to closest monochrome color with a 10-20 offset)
4. Data binding in accordions is functioning one-way (sliders / entries not getting updated, data inside structs - does)
5. Data update within the main struct is working by comparing its copy to itself - should redo it to use more efficient way (unreleased version at least uses a timer... still terribly inefficient however)
