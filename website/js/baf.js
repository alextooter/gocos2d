/*
 *
 * Buttons and Forms
 *
 * Repo: https://github.com/SimonWaldherr/buttons-and-forms
 * Demo: http://simonwaldherr.github.com/buttons-and-forms/demo/
 * Editor: http://simonwaldherr.github.com/buttons-and-forms/editor/
 * License: MIT
 * Version: 1.7
 *
 */

function $id(id)
  {
    return document.getElementById(id);
  }

var baf_keepCalling = true;

function baf_plusone(id, wait, min, slider)
  {
    if((baf_keepCalling)&&((Math.round($id(id).max)) > $id(id).value))
      {
        $id(id).value++;
        if(wait > min)
          {
            wait = wait-60;
          }
        if(wait < min)
          {
            wait = min;
          }
        if(slider == 1)
          {
            fdSlider.updateSlider(id);
          }
        window.setTimeout(baf_plusone, wait, id, wait, min, slider);
      }
  }

function baf_minusone(id, wait, min, slider)
  {
    if((baf_keepCalling)&&($id(id).value > (Math.round($id(id).min))))
      {
        $id(id).value--;
        if(wait > min)
          {
            wait = wait-60;
          }
        if(wait < min)
          {
            wait = min;
          }
        if(slider == 1)
          {
            fdSlider.updateSlider(id);
          }
        window.setTimeout(baf_minusone, wait, id, wait, min, slider);
      }
  }

function baf_changeValues(element)
  {
    
    if(element != '[object HTMLInputElement]')
      {
        element = $id(element);
      }
    var kc = window.event.keyCode;
    if (((kc === 37)||(kc === 40)||(kc === 109))&&(element.value > (Math.round(element.min))))
      {
        //left||down||minus -
        element.value = Math.round(element.value)-1;
        return false;
      }
    else if (((kc === 38)||(kc === 39)||(kc === 107))&&((Math.round(element.max)) > element.value))
      {
        //up||right||plus +
        element.value = Math.round(element.value)+1;
        return false;
      }
    else if ((kc === 34)&&(element.value > (Math.round(element.min))))
      {
        //page down -
        element.value = Math.round(element.value)-10;
        return false;
      }
    else if ((kc === 33)&&((Math.round(element.max)) > element.value))
      {
        //page up +
        element.value = Math.round(element.value)+10;
        return false;
      }
    else
      {
        return true;
      }
  }

function baf_changeloadingmode(element)
  {
    if(element.disabled == true)
      {
        
      }
    else
      {
        if(element.className.search('loading') != -1)
          {
            element.className = element.className.replace(' loading', '');
          }
        else
          {
            element.className = element.className+' loading';
            return true;
          }
      }
  }

function baf_dropdown()
  {
    this.offsetParent.className = this.offsetParent.className+" open";
    window.setTimeout(baf_eventddclose,10);
  }

function baf_eventddclose()
  {
    var html = document.getElementsByTagName("body"); 
    html[0].addEventListener("click", baf_dropdownclose, false);
  }

function baf_dropdownclose()
  {
    var bafele   = new Array();
    bafele['dd'] = document.getElementsByTagName('*');
    for(x in bafele['dd'])
    {
      classn = bafele['dd'][x].className;
      if((typeof classn == 'string'))
        {
          if(classn.length > 3)
            {
              if(classn.search("dropdown-toggle") != -1)
                {
                  if(bafele['dd'][x].offsetParent.className.search("open") != -1)
                    {
                      bafele['dd'][x].offsetParent.className = bafele['dd'][x].offsetParent.className.replace('open', '');
                    }
                }
            }
        }
    }
    var html = document.getElementsByTagName("body"); 
    html[0].removeEventListener("click", baf_dropdownclose, false);
  }

function baf_listenerInit()
  {
    var classn, maxvalue, minvalue;
    var bafele   = new Array();
    var rangeSlider = new Array();
    bafele['dd'] = document.getElementsByTagName('*');
    for(i in bafele['dd'])
      {
        classn = bafele['dd'][i].className;
        if((typeof classn == 'string'))
          {
            if(classn.length > 3)
              {
                if(classn.search("dropdown-toggle") != -1)
                  {
                    bafele['dd'][i].addEventListener("click", baf_dropdown, false);
                  }
                if((classn.search("range-slider") != -1)&&(bafele['dd'][i].tagName.toLowerCase() == 'input'))
                  {
                    rangeSlider[i] = bafele['dd'][i].id;
                  }
              }
          }
      }
    for(i in rangeSlider)
      {
        if(typeof Math.round($id(rangeSlider[i]).getAttribute("max")) != 'number')
          {
            maxvalue = 100;
          }
        else
          {
            maxvalue = Math.round($id(rangeSlider[i]).getAttribute("max"));
          }
        if(typeof Math.round($id(rangeSlider[i]).getAttribute("min")) != 'number')
          {
            minvalue = 0;
          }
        else
          {
            minvalue = Math.round($id(rangeSlider[i]).getAttribute("min"));
          }
        fdSlider.createSlider({
          inp:$id(rangeSlider[i]),
          step:1, 
          maxStep:1,
          min:minvalue,
          max:maxvalue,
          animation:"tween",
          forceValue:true
        });
      }
  }



function baf_dgeb(name, type)
  {
    if (type == 'name')
      {
        if (!document.getElementsByTagName(name))
          {
            return false;
          }
        return document.getElementsByTagName(name);
      }
    else
    {
      if (!$id(name))
        {
          return false;
        }
      return $id(name);
    }
  }

function baf_colorize(data,output)
  {
    if (typeof(data)=='number')
      {
        if ((data < 115)&&(data > 1))
          {
            boxcolor = 'rgb(255,'+(153+data)+','+(153-data)+')';
            baf_dgeb(output, 'id').style.background = boxcolor;
          }
        if ((data > 115)&&(data < 230))
          {
            data = data - 115;
            boxcolor = 'rgb('+(255-data)+',243,63)';
            baf_dgeb(output, 'id').style.background = boxcolor;
          }
        if (data > 230)
          {
            boxcolor = 'rgb(145,243,63)';
            baf_dgeb(output, 'id').style.background = boxcolor;
          }
      }
    else if (data == 'none')
      {
        baf_dgeb(output, 'id').style.background = 'rgb(204,204,204)';
      }
    else if (data == true)
      {
        if (baf_dgeb(output, 'id').style.background = 'rgb(145,243,63)')
          {
            return true;
          }
        return false;
      }
    else
      {
        if (baf_dgeb(output, 'id').style.background = 'rgb(255,153,153)')
          {
            return true;
          }
        return false;
      }
  }

function baf_converttxt(text)
  {
    var textarray = text.split("");
    var textoutput = "";
    for(var i in textarray)
      {
        //alert(textarray[i].charCodeAt(0));
        if((textarray[i].charCodeAt(0)>47)&&(textarray[i].charCodeAt(0)<59))
          {
            textoutput = textoutput+textarray[i]; //0-9, :
          }
        if((textarray[i].charCodeAt(0)>62)&&(textarray[i].charCodeAt(0)<91))
          {
            textoutput = textoutput+textarray[i]; //A-Z
          }
        if((textarray[i].charCodeAt(0)>96)&&(textarray[i].charCodeAt(0)<123))
          {
            textoutput = textoutput+textarray[i]; //a-z
          }
        if((((textarray[i].charCodeAt(0)>34)&&(textarray[i].charCodeAt(0)<47))||(textarray[i].charCodeAt(0)==61)||(textarray[i].charCodeAt(0)==33))&&(textarray[i].charCodeAt(0)!=39))
          {
            textoutput = textoutput+textarray[i]; //#$%&()*+,-.=!
          }
      }
    return textoutput;
  }

function baf_convertnumber(text)
  {
    var textarray = text.split("");
    var textoutput = "";
    var dots = 0;
    for(var i in textarray)
      {
        //alert(textarray[i].charCodeAt(0));
        if((textarray[i].charCodeAt(0)>47)&&(textarray[i].charCodeAt(0)<59))
          {
            textoutput = textoutput+textarray[i]; //0-9, :
          }
        if(((textarray[i].charCodeAt(0)==46)||(textarray[i].charCodeAt(0)==44))&&(dots == 0))
          {
            textoutput = textoutput+'.'; //#$%&()*+,-.=!
            dots = 1;
          }
      }
    return textoutput;
  }

function baf_age(dateid,notation,rule,output)
  {
    var date = baf_dgeb(dateid, 'id').value;
    var splited = "";
    var day = 0;
    var month = 0;
    var year = 0;


    switch(notation)
      {
        case "dd.mm.yyyy":
          splited = date.split(".");
          day = splited[0];
          month = splited[1];
          year = splited[2];
          break;
        case "yyyy.mm.dd":
          splited = date.split(".");
          day = splited[2];
          month = splited[1];
          year = splited[0];
          break;
        case "mm/dd/yyyy":
          splited = date.split("/");
          day = splited[1];
          month = splited[0];
          year = splited[2];
          break;
        default:
          break;
      }

    var schaltjahr = 0;
    if(((year % 4 == 0) && ((year % 100!= 0) || (year % 400 == 0))))
      {
        schaltjahr = 1;
      }

    var error = 0;
    if((day<1)||(day>31)||(month<1)||(month>12))
      {
        error = 1;
      }
    if(((month==4)||(month==6)||(month==9)||(month==11))&&(day>30))
      {
        error = 1;
      }
    if(month==2)
      {
        if((schaltjahr==1)&&(day>29))
          {
            error = 1;
          }
        else if((schaltjahr==0)&&(day>28))
          {
            error = 1;
          }
      }

    var now = new Date();
    var nowts = now.getTime()/31556952000;  
    var indate = new Date(year, month, day, 1,0,0);
    var timestamp = indate.getTime()/31556952000;

    if((year<1000)||(year>now.getFullYear()))
      {
        error = 1;
      }

    if((isNaN(timestamp))||(error == 1))
      {
        baf_colorize(false,output);
        return false;
      }
    else
      {

        baf_number(nowts-timestamp,rule,output);
        return true;
      }
  }

function baf_date(dateid,notation,output)
  {
    var date = baf_dgeb(dateid, 'id').value;
    var splited = "";
    var day = 0;
    var month = 0;
    var year = 0;

    switch(notation)
      {
        case "dd.mm.yyyy":
          splited = date.split(".");
          day = splited[0];
          month = splited[1];
          year = splited[2];
          break;
        case "yyyy.mm.dd":
          splited = date.split(".");
          day = splited[2];
          month = splited[1];
          year = splited[0];
          break;
        case "mm/dd/yyyy":
          splited = date.split("/");
          day = splited[1];
          month = splited[0];
          year = splited[2];
          break;
        default:
          break;
      }

    var schaltjahr = 0;
    if(((year % 4 == 0) && ((year % 100!= 0) || (year % 400 == 0))))
      {
        schaltjahr = 1;
      }

    var error = 0;
    if((day<1)||(day>31)||(month<1)||(month>12))
      {
        error = 1;
      }
    if(((month==4)||(month==6)||(month==9)||(month==11))&&(day>30))
      {
        error = 1;
      }
    if(month==2)
      {
        if((schaltjahr==1)&&(day>29))
          {
            error = 1;
          }
        else if((schaltjahr==0)&&(day>28))
          {
            error = 1;
          }
      }

    var indate = new Date(year, month, day, 1,0,0);
    var timestamp = indate.getTime()/1000;


    if((isNaN(timestamp))||(error == 1))
      {
        baf_colorize(false,output);
        return false;
      }
    else
      {
        baf_colorize(true,output);
        return true;
      }
  }

function baf_email(email,output)
  {
    var reg = /^([A-Za-z0-9_\-\.])+\@([A-Za-z0-9_\-\.])+\.([A-Za-z]{2,7})$/;
    if(reg.test(email) == false)
      {
        baf_colorize(false,output);
        return false;
      }
    else
      {
        baf_colorize(true,output);
        return true;
      }
  }

function baf_number(innumber,rule,output)
  {
    var checkf = 0;
    var checks = 0;
    var number = parseInt(innumber);
    var rulenumbers = rule.split("-");

    if((number>=rulenumbers[0])||(rulenumbers[0]=='x'))
      {
        checkf = 1;
      }
    if((number<=rulenumbers[1])||(rulenumbers[1]=='x'))
      {
        checks = 1;
      }
    if(isNaN(innumber))
      {
        checkf = 0;
      }

    if((checkf == 1)&&(checks == 1))
      {
        baf_colorize(true,output);
        return true;
      }
    else
      {
        baf_colorize(false,output);
        return false;
      }
  }

function baf_password(password,output)
  {
    var valunicode;
    var keys = password.split("");
    var numbers = 1;
    var uletter = 1;
    var lletter = 1;
    var special = 1;
    var complex = 0;
    var boxcolor = '';
    for(var i = 0; i < keys.length; i++)
      {
        valunicode = keys[i].charCodeAt(0);
        if((valunicode > 0x40)&&(valunicode < 0x5B)) //Großbuchstaben A-Z
          {
            ++uletter;
          }
        else if((valunicode > 0x60)&&(valunicode < 0x7B)) //Kleinbuchstaben a-z
          {
            ++lletter;
          }
        else if((valunicode > 0x2F)&&(valunicode < 0x3A)) //Zahlen 0-9
          {
            ++numbers;
          }
        else if((valunicode > 0x20)&&(valunicode < 0x7F)) // Sonderzeichen
          {
            ++special;
          }
        else if((valunicode < 0x21)||(valunicode > 0x7E))
          {

          }
      }
    complex = ((uletter*lletter*numbers*special)+Math.round(uletter*1.8+lletter*1.5+numbers+special*2))-6;
    baf_colorize(complex,output);
    return complex;
  }

function baf_repeat(idone,idtwo,output)
  {
    if(baf_dgeb(idone, 'id').value != baf_dgeb(idtwo, 'id').value)
      {
        baf_colorize(false,output);
        return false;
      }
    else
      {
        baf_colorize(true,output);
        return true;
      }
  }

function baf_cleartext(text,output)
  {
    var textoutput = baf_converttxt(text);
    baf_dgeb(output, 'id').value = textoutput;
  }

function baf_clearfloat(text,output)
  {
    var textoutput = baf_convertnumber(text);
    baf_dgeb(output, 'id').value = textoutput;
  }
