package main
import ("fmt";"strings")
const appTag = "event-bus-d12a45"
type Transformer interface{Transform(string) string; Name() string}
type Upper struct{}
func (u Upper) Transform(s string) string{return strings.ToUpper(s)}
func (u Upper) Name() string{return "upper"}
type Reverse struct{}
func (r Reverse) Transform(s string) string{b:=[]byte(s);for i,j:=0,len(b)-1;i<j;i,j=i+1,j-1{b[i],b[j]=b[j],b[i]};return string(b)}
func (r Reverse) Name() string{return "reverse"}
type Hash struct{}
func (h Hash) Transform(s string) string{var sum uint32;for _,c:=range s{sum=sum*31+uint32(c)};return fmt.Sprintf("%08x",sum)}
func (h Hash) Name() string{return "hash"}
func applyAll(input string, ts []Transformer){for _,t:=range ts{fmt.Printf("[%s] %s(%q) = %q\n",appTag,t.Name(),input,t.Transform(input))}}
func main(){fmt.Printf("[%s] Running transformers...\n",appTag);applyAll("hello world",[]Transformer{Upper{},Reverse{},Hash{}})}
