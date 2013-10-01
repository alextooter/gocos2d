package gocos2d

import gl "github.com/mortdeus/egles/es2"
import "fmt"

const (
	SHADER_POSITION_COLOR_FRAG = `

	precision lowp float;                 
        	
        varying vec4 v_fragmentColor;         
                                              
        void main(){                                     
            	gl_FragColor = v_fragmentColor;       
        }`

	SHADER_POSITION_COLOR_VERT = `

       	attribute vec4 a_position;                
       	attribute vec4 a_color;                                                             
       	
       	varying lowp vec4 v_fragmentColor;        
       	                                          
       	void main(){                                         
       	    	gl_Position = CC_MVPMatrix * a_position;  
       	    	gl_Position = (CC_PMatrix * CC_MVMatrix) * a_position;  
       	    	v_fragmentColor = a_color;             
       	}`

	SHADER_POSITION_UCOLOR_FRAG = `                                           

       	precision lowp float;                    
       	
       	varying vec4 v_fragmentColor;            
                                                
       	void main(){                                        
           	gl_FragColor = v_fragmentColor;      
       	}`

	SHADER_POSITION_UCOLOR_VERT = `                                    

        attribute vec4 a_position;               
        uniform vec4 u_color;                 
        uniform float u_pointSize;               
                                                 
        varying lowp vec4 v_fragmentColor;       
                                                 
        void main(){                                        
        	gl_Position = CC_MVPMatrix * a_position;  
        	gl_Position = (CC_PMatrix * CC_MVMatrix) * a_position;  
        	gl_PointSize = u_pointSize;          
        	v_fragmentColor = u_color;           
        }`

	SHADER_POSITION_COLOR_LENGTH_TEXTURE_FRAG = `

        // #extension GL_OES_standard_derivatives : enable
       
        varying mediump vec4 v_color;																								
        varying mediump vec2 v_texcoord;																							    
       
        void main(){																														
       
        // #if defined GL_OES_standard_derivatives																					
        // gl_FragColor = v_color*smoothstep(0.0, length(fwidth(v_texcoord)), 1.0 - length(v_texcoord));							    
        // #else																														
       
        	gl_FragColor = v_color * step(0.0, 1.0 - length(v_texcoord));														        
        // #endif																													
        }`

	SHADER_POSITION_COLOR_LENGTH_TEXTURE_VERT = `

        attribute mediump vec4 a_position;
        attribute mediump vec2 a_texcoord;
        attribute mediump vec4 a_color;
       
        varying mediump vec4 v_color;
        varying mediump vec2 v_texcoord;
        
        void main(){
        	v_color = a_color;
        	
        	//vec4(a_color.rgb * a_color.a, a_color.a);	
        	
        	v_texcoord = a_texcoord;																						      
         	
         	gl_Position = CC_MVPMatrix * a_position;
            	gl_Position = (CC_PMatrix * CC_MVMatrix) * a_position;
        }`

	SHADER_POSITION_TEXTURE_FRAG = `

        precision lowp float;                      
                                                  
        varying vec2 v_texCoord;                   
        uniform sampler2D CC_Texture0;             
                                                  
        void main(){                                          
            	gl_FragColor =  texture2D(CC_Texture0, v_texCoord);   
        }`

	SHADER_POSITION_TEXTURE_VERT = `

        attribute vec4 a_position;                   
        attribute vec2 a_texCoord;                  
                                                    
        varying mediump vec2 v_texCoord;           
                                                   
        void main(){                                          
        	//gl_Position = CC_MVPMatrix * a_position;  
            	
            	gl_Position = (CC_PMatrix * CC_MVMatrix) * a_position;  
            	v_texCoord = a_texCoord;               
        }`

	SHADER_POSITION_TEXTURE_UCOLOR_FRAG = `                                                

        precision lowp float;                        
                                                     
        uniform vec4 u_color;                        
        varying vec2 v_texCoord;                     
                                                     
        uniform sampler2D CC_Texture0;               
                                                     
        void main(){                                            
            	gl_FragColor =  texture2D(CC_Texture0, v_texCoord) * u_color;    
        }`

	SHADER_POSITION_TEXTURE_UCOLOR_VERT = `

        attribute vec4 a_position;                   
        attribute vec2 a_texCoord;                   
                                                     
        varying mediump vec2 v_texCoord;             
                                                     
        void main(){                                            
        	//gl_Position = CC_MVPMatrix * a_position;  
            	
            	gl_Position = (CC_PMatrix * CC_MVMatrix) * a_position;  
            	v_texCoord = a_texCoord;                 
        }`

	SHADER_POSITION_TEXTURE_A8COLOR_FRAG = `

        precision lowp float;                        
                                                     
        varying vec4 v_fragmentColor;                
        varying vec2 v_texCoord;                     
        uniform sampler2D CC_Texture0;                 
                                                     
        void main(){                                            
            	gl_FragColor = vec4(
            		// RGB from uniform
            		v_fragmentColor.rgb,      
            			
            		// A from texture and uniform                              
                	v_fragmentColor.a * texture2D(CC_Texture0, v_texCoord).a
            	);                                         
        }`

	SHADER_POSITION_TEXTURE_A8COLOR_VERT = `

        attribute vec4 a_position;                   
        attribute vec2 a_texCoord;                   
        attribute vec4 a_color;                      
                                                     
        varying lowp vec4 v_fragmentColor;           
        varying mediump vec2 v_texCoord;             
                                                     
        void main(){                                            
        	//gl_Position = CC_MVPMatrix * a_position;  
            	
            	gl_Position = (CC_PMatrix * CC_MVMatrix) * a_position;  
            	v_fragmentColor = a_color;               
            	v_texCoord = a_texCoord;                 
        }`

	SHADER_POSITION_TEXTURE_COLOR_FRAG = `

        precision lowp float;                        
                                                     
        varying vec4 v_fragmentColor;                
        varying vec2 v_texCoord;                     
        uniform sampler2D CC_Texture0;               
                                                     
        void main(){                                            
            	gl_FragColor = v_fragmentColor * texture2D(CC_Texture0, v_texCoord);         
        }`

	SHADER_POSITION_TEXTURE_COLOR_VERT = `

        attribute vec4 a_position;                   
        attribute vec2 a_texCoord;                   
        attribute vec4 a_color;                      
                                                     
        varying lowp vec4 v_fragmentColor;           
        varying mediump vec2 v_texCoord;             
                                                     
        void main(){                                            
        	//gl_Position = CC_MVPMatrix * a_position;  
            	
            	gl_Position = (CC_PMatrix * CC_MVMatrix) * a_position;  
            	v_fragmentColor = a_color;               
            	v_texCoord = a_texCoord;                 
        }`

	SHADER_POSITION_TEXTURE_COLOR_ALPHATEST_FRAG = `

        precision lowp float;                           
                                                        
        varying vec4 v_fragmentColor;                   
        varying vec2 v_texCoord;                        
        uniform sampler2D CC_Texture0;                  
        uniform float CC_alpha_value;                   
                                                        
        void main(){                                               
            	vec4 texColor = texture2D(CC_Texture0, v_texCoord);          
                                                        
            	/* 
        	mimic: glAlphaFunc(GL_GREATER)           
        	pass if ( incoming_pixel >= CC_alpha_value ) => fail if incoming_pixel < CC_alpha_value         
                */                                        
            	if ( texColor.a <= CC_alpha_value )          
                discard;                                
                                                        
            	gl_FragColor = texColor * v_fragmentColor;  
        }`

	SHADEREX_SWITCHMASK_FRAG = `

        precision lowp float;                            
                                                         
        varying vec4 v_fragmentColor;                    
        varying vec2 v_texCoord;                         
        uniform sampler2D u_texture;                     
        uniform sampler2D   u_mask;                      
                                                         
        void main(){                                                
            vec4 texColor   = texture2D(u_texture, v_texCoord);          
            vec4 maskColor  = texture2D(u_mask, v_texCoord);             
            vec4 finalColor = vec4(texColor.r, texColor.g, texColor.b, maskColor.a * texColor.a);        
            gl_FragColor    = v_fragmentColor * finalColor;              
        }`
)

func FragmentShader(s string) uint {
	shader := gl.CreateShader(gl.FRAGMENT_SHADER)
	gl.ShaderSource(shader, s)
	gl.CompileShader(shader)

	if gl.GetShaderiv(shader, gl.COMPILE_STATUS, make([]int32, 1))[0] == 0 {
		fmt.Printf("Fragment Shader:\n%s\n", gl.GetShaderInfoLog(shader, 1000))

	}
	return shader
}
func VertexShader(s string) uint {
	shader := gl.CreateShader(gl.VERTEX_SHADER)
	gl.ShaderSource(shader, s)
	gl.CompileShader(shader)

	if gl.GetShaderiv(shader, gl.COMPILE_STATUS, make([]int32, 1))[0] == 0 {
		fmt.Printf("Vertex Shader:\n%s\n", gl.GetShaderInfoLog(shader, 1000))
	}
	return shader
}
func Program(fsh, vsh uint) uint {
	p := gl.CreateProgram()
	gl.AttachShader(p, fsh)
	gl.AttachShader(p, vsh)
	gl.LinkProgram(p)
	if gl.GetProgramiv(p, gl.LINK_STATUS, make([]int32, 1))[0] == 0 {
		fmt.Printf("Program:\n%s\n", gl.GetProgramInfoLog(p, 1000))
	}
	return p
}
