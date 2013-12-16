package gocos2d

import gl "github.com/mortdeus/egles/es2"
import "fmt"

type Shader interface {
	GetShader(uint) uint
	SetShader(uint, uint)
}

func NewShader(s string, typ uint) (shader uint) {

	switch typ {
	case gl.FRAGMENT_SHADER:
		shader = gl.CreateShader(gl.FRAGMENT_SHADER)
		gl.ShaderSource(shader, s)
		gl.CompileShader(shader)
		if gl.GetShaderiv(shader, gl.COMPILE_STATUS, make([]int32, 1))[0] == 0 {
			fmt.Printf("Frag:\n%s\n", gl.GetShaderInfoLog(shader, 1000))
		}
	case gl.VERTEX_SHADER:
		shader = gl.CreateShader(gl.VERTEX_SHADER)
		gl.ShaderSource(shader, s)
		gl.CompileShader(shader)
		if gl.GetShaderiv(shader, gl.COMPILE_STATUS, make([]int32, 1))[0] == 0 {
			fmt.Printf("Vert:\n%s\n", gl.GetShaderInfoLog(shader, 1000))
		}
	}
	return shader

}
func Program(vsh, fsh uint) uint {
	p := gl.CreateProgram()
	gl.AttachShader(p, fsh)
	gl.AttachShader(p, vsh)
	gl.LinkProgram(p)
	if gl.GetProgramiv(p, gl.LINK_STATUS, make([]int32, 1))[0] == 0 {
		fmt.Printf("Program:\n%s\n", gl.GetProgramInfoLog(p, 1000))
	}
	return p
}

const (
	POSITION_COLOR_FRAG = `

	    precision lowp float;                      	

        varying vec4 v_fragColor;         
                                          
        void main(){                                     
            gl_FragColor = v_fragColor;       
        }
    `

	POSITION_COLOR_VERT = `

	    uniform mat4 MVPMatrix;

       	attribute vec4 a_pos;                
       	attribute vec4 a_color;

       	varying lowp vec4 v_fragColor;        
       	                                          
       	void main(){                                         
       	    	gl_Position = MVPMatrix * a_pos;   
       	    	v_fragColor = a_color;             
       	}
    `

	POSITION_UCOLOR_FRAG = `                                           

       	precision lowp float;                        	

       	varying vec4 v_fragColor;            
                                                
       	void main(){                                        
           	gl_FragColor = v_fragColor;      
       	}
    `

	POSITION_UCOLOR_VERT = `                                    

        attribute vec4 a_pos;                    

        uniform vec4 u_color;                 
        uniform float u_pointSize;               
        uniform mat4 MVPMatrix;

        varying lowp vec4 v_fragColor;       
                                                 
        void main(){                                        
        	gl_Position = MVPMatrix * a_pos;    
        	gl_PointSize = u_pointSize;          
        	v_fragColor = u_color;           
        }
    `

	POSITION_COLOR_LENGTH_TEXTURE_FRAG = `
       
        varying mediump vec4 v_color;
        varying mediump vec2 v_texcoord;

        void main(){
        	gl_FragColor = v_color * step(0.0, 1.0 - length(v_texcoord));
        }
    `

	POSITION_COLOR_LENGTH_TEXTURE_VERT = `

	    uniform mat4 MVPMatrix;

        attribute mediump vec4 a_pos;
        attribute mediump vec2 a_texcoord;
        attribute mediump vec4 a_color;
       
        varying mediump vec4 v_color;
        varying mediump vec2 v_texcoord;
        
        void main(){
        	v_color = a_color;	

        	//vec4(a_color.rgb * a_color.a, a_color.a);	 	

        	v_texcoord = a_texcoord;
         	gl_Position = MVPMatrix * a_pos;
            	
        }`

	POSITION_TEXTURE_FRAG = `

        precision lowp float;                      
                                                  
        varying vec2 v_texCoord;                   
        uniform sampler2D Texture0;             
                                                  
        void main(){                                          
            	gl_FragColor =  texture2D(Texture0, v_texCoord);   
        }
    `

	POSITION_TEXTURE_VERT = `

	    uniform mat4 MVPMatrix;

        attribute vec4 a_pos;                   
        attribute vec2 a_texCoord;                  
                                                    
        varying mediump vec2 v_texCoord;           
                                                   
        void main(){                                          
            gl_Position = MVPMatrix * a_pos;  
            v_texCoord = a_texCoord;               
        }
    `

	POSITION_TEXTURE_UCOLOR_FRAG = `                                                

        precision lowp float;                        
                                                     
        uniform vec4 u_color;                        
        uniform sampler2D CC_Texture0; 

        varying vec2 v_texCoord;                     
                                                                                                   
        void main(){                                            
            	gl_FragColor =  texture2D(CC_Texture0, v_texCoord) * u_color;    
        }
    `

	POSITION_TEXTURE_UCOLOR_VERT = `
	
	    uniform mat4 MVPMatrix;

        attribute vec4 a_pos;                   
        attribute vec2 a_texCoord;                   
                                                     
        varying mediump vec2 v_texCoord;             
                                                     
        void main(){                                            
        	gl_Position = MVPMatrix * a_pos;  
            v_texCoord = a_texCoord;                 
        }
    `

	POSITION_TEXTURE_A8COLOR_FRAG = `

        precision lowp float;                        
                                                     
        varying vec4 v_fragColor;                
        varying vec2 v_texCoord;                     
        uniform sampler2D CC_Texture0;                 
                                                     
        void main(){                                            
            gl_FragColor = vec4(
                // RGB from uniform   
                v_fragColor.rgb,

                // A from texture and uniform      
                v_fragColor.a * texture2D(CC_Texture0, v_texCoord).a);                                         
        }
    `

	POSITION_TEXTURE_A8COLOR_VERT = `

	    uniform mat4 MVPMatrix;

        attribute vec4 a_pos;                   
        attribute vec2 a_texCoord;                   
        attribute vec4 a_color;                      
                                                     
        varying lowp vec4 v_fragColor;           
        varying mediump vec2 v_texCoord;             
                                                     
        void main(){                                            
        	gl_Position = MVPMatrix * a_pos;  	  
            v_fragColor = a_color;               
            v_texCoord = a_texCoord;                 
        }
    `

	POSITION_TEXTURE_COLOR_FRAG = `

        precision lowp float;                        
                                                     
        varying vec4 v_fragColor;                
        varying vec2 v_texCoord;                     
        uniform sampler2D Tex0;               
                                                     
        void main(){                                            
            gl_FragColor = v_fragColor * texture2D(Tex0, v_texCoord);         
        }
    `

	POSITION_TEXTURE_COLOR_VERT = `

	    uniform mat4 MVPMatrix;

        attribute vec4 a_pos;                   
        attribute vec2 a_texCoord;                   
        attribute vec4 a_color;                      
                                                     
        varying lowp vec4 v_fragColor;           
        varying mediump vec2 v_texCoord;             
                                                     
        void main(){                                            
        	gl_Position = MVPMatrix * a_pos;  
            v_fragColor = a_color;               
            v_texCoord = a_texCoord;                 
        }
    `

	POSITION_TEXTURE_COLOR_ALPHATEST_FRAG = `

        precision lowp float;                           
                                                        
        varying vec4 v_fragColor;                   
        varying vec2 v_texCoord;                        
        uniform sampler2D Tex0;                  
        uniform float alpha_value;                   
                                                        
        void main(){                                               
            vec4 texColor = texture2D(Tex0, v_texCoord);          
                                                        
            // mimic: glAlphaFunc(GL_GREATER)           
        	
            // pass if ( incoming_pixel >= alpha_value ) 
		    // => fail if incoming_pixel < alpha_value                                                
            
            if ( texColor.a <= alpha_value ) discard;

            gl_FragColor = texColor * v_fragColor;  
        }
    `

	EX_SWITCHMASK_FRAG = `

        precision lowp float;                            
                                                         
        varying vec4 v_fragColor;                    
        varying vec2 v_texCoord;                         
        uniform sampler2D u_texture;                     
        uniform sampler2D u_mask;                      
                                                         
        void main(){                                                
            vec4 texColor   = texture2D(u_texture, v_texCoord);          
            vec4 maskColor  = texture2D(u_mask, v_texCoord);             
            vec4 finalColor = vec4(texColor.r, texColor.g, texColor.b, maskColor.a * texColor.a);        
            gl_FragColor    = v_fragColor * finalColor;              
        }
    `
)
